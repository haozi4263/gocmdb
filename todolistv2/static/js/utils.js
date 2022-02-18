
// 对 < > & ' "进行替换，用于解决xsrf攻击。在所有html显示对地方使用htmlEncode对字符串进行编码
function htmlEncode(str){
    if(typeof(str) == "undefined"){
        return "";
    }
    if(typeof(str) != "string"){
        str = str.toString();
    }
    return str.replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/&/g, '&amp;').replace(/'/g,'&#39;').replace(/"/g,'&quot;');
}