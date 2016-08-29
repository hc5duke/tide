var ID3 = require('id3-parser');
var fs = require('fs');

var filePath = './sample.mp3';
var fileBuffer = fs.readFileSync(filePath);

ID3.parse(fileBuffer).then(function(tag) {
    console.log(tag.image.data.length);
});
