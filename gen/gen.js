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
    "unsigned long": "int",
    "long": "int",
    "boolean": "bool",
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

    _.map(tree, function(iface) {
        if (iface.type != "interface") {
            return;
        }

        if (iface.partial) {
            if (ifaces[iface.name]) {
                ifaces[iface.name].members.concat(iface.members);
            }

            return;
        }

        iface.tagType = mapTag(iface.name);
        if (!iface.tagType) {
            return;
        }

        iface.htmlTag = iface.tagType.toLowerCase();
        iface.privType = "HTML" + iface.tagType;
        ifaces[iface.name] = iface
    })

    function parentRecursive(iface, fn, pp=".") {
        fn(iface, pp);
        if (iface.parentIface) {
            parentRecursive(iface.parentIface, fn, pp + "p.");      
        }
    }

    _.forOwn(ifaces, function(iface) {
        iface.parentIface = null;
        if (iface.inheritance) {
            iface.parentIface = ifaces[iface.inheritance];
        }
    })

    _.forOwn(ifaces, function(iface) {
        if (iface.parentIface &&
                iface.parentIface.tagType == "Element" && iface.members.length == 0) {

            emptyTags.push(iface);
            return;
        }

        iface._ = _;
        iface.goTypeMap = goTypeMap;
        iface.parentRecursive = parentRecursive;
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
