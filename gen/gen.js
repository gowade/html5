var _ = require('lodash');
var fs = require("fs")
var WebIDL2 = require("webidl2");
var ejs = require('ejs');

var template = fs.readFileSync("gen/template.go.in", {encoding: "utf8"})
var emptyTagsTemp = fs.readFileSync("gen/emptytags.go.in", {encoding: "utf8"})

var tagMap = {
    "UList": "UL",
    "OList": "OL",
    "Anchor": "A",
    "": "Element",
}

var goTypeMap = {
    "DOMString": "string", 
    "USVString": "string", 
    "unsigned long": "int",
    "long": "int",
    "boolean": "bool",
}

function toGoType(typeName) {
    if (_.endsWith(typeName, "EventHandler")) {
        return "Action";
    }
    return goTypeMap[typeName]
}

function mapTag(name) {
    var match = name.match(/^HTML(.*)Element$/);
    if (!match) {
        return null;
    }

    var mtag = match[1];
    if (mtag == "Unknown") {
        return null;
    }

    return tagMap[mtag] ? tagMap[mtag] : mtag;
}

function parse(file) {
    var tree = WebIDL2.parse(fs.readFileSync(file, {encoding: "utf8"}));
    var ifaces = {};
    var emptyTags = [];

    _.each(tree, function(item) {
        if (item.type == "implements") {
            if (ifaces[item.target]) {
                if (!ifaces[item.target].impls) {
                    ifaces[item.target].impls = []
                }
                ifaces[item.target].impls.push(item.implements);
            }
            return;
        }

        if (item.type != "interface") {
            return;
        }

        iface = item

        if (iface.partial) {
            if (ifaces[iface.name]) {
                ifaces[iface.name].members.concat(iface.members);
            }

            return;
        }

        iface.tagType = mapTag(iface.name);
        if (iface.tagType) {
            iface.htmlTag = iface.tagType.toLowerCase();
            iface.privType = "HTML" + iface.tagType;
        }

        ifaces[iface.name] = iface
    })

    function memberMap(members) {
        return _.zipObject(_.map(members, "name"), members)
    }

    _.forOwn(ifaces, function(iface) {
        if (!iface.tagType) {
            return;
        }

        iface.allMembers = memberMap(iface.members);
        iface.parentIface = null;
        if (iface.inheritance) {
            iface.parentIface = ifaces[iface.inheritance];
            _.defaults(iface.allMembers, memberMap(iface.parentIface.members));
        }
    
        _.each(iface.impls, function(implName) {
            var im = ifaces[implName];
            if (!im) {
                return;
            }

            if (_.includes(implName, "Event") && iface.tagType != "Element") {
                return;
            }

            _.defaults(iface.allMembers, memberMap(im.members));
        })
    })

    _.forOwn(ifaces, function(iface) {
        if (!iface.tagType) {
            return;
        }

        if (iface.parentIface &&
                iface.parentIface.tagType == "Element" && iface.members.length == 0) {

            emptyTags.push(iface);
            return;
        }

        iface._ = _;
        iface.toGoType = toGoType;
        iface.toGoMethodName = function(name) {
            switch (name) {
            case "id":
                return "ID"
            break;
            case "htmlFor":
                return "For"
            }
            return _.upperFirst(name);
        }

        var output = ejs.render(template, iface);
        var fileName = "html_" + iface.tagType.toLowerCase() + ".go";
        fs.writeFile(fileName, output, function(err) {
            if(err) {
                return console.log(err);
            }
        });
    })

    return emptyTags;
}

var emptyTags = parse("gen/idl/html5.idl");

(function() {
    var output = ejs.render(emptyTagsTemp, {
        tags: emptyTags,
        _: _,
    });

    var fileName = "html_emptytags.go";
    fs.writeFile(fileName, output, function(err) {
        if(err) {
            return console.log(err);
        }
    });
})();

console.log("code generation complete.")
