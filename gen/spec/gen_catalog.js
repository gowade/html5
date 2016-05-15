var jq = require('cheerio');
var fs = require('fs');


var doc = fs.readFileSync("gen/spec/catalog_src.html", {encoding: 'utf8'});
$ = jq.load(doc);

var list = [];

$("a").each(function(_i, a) {
    var href = $(a).attr("href");
    if (!href) {
        return;
    }

    var match = href.match(/#the-([a-z]+)-element/);
    if (!match) {
        return; 
    }

    list.push(match[1]);
})

fs.writeFile("gen/spec/catalog.json", JSON.stringify(list), function(err) {
    if(err) {
        return console.log(err);
    }
});
