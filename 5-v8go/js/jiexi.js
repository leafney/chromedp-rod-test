
function Jiexi(html) {
    const $ = cheerio.load(html);

    const text= $(".title").text();
    console.log(text);
    return text;
}