const CryptoJS = require('crypto-js');

var date = new Date();
var yyyy = date .getUTCFullYear().toString();
var mm = (date.getUTCMonth() + 1).toString();
var dd = date.getUTCDate().toString();
var today = yyyy + '-' +(mm[1] ? mm : "0"+mm[0])+ '-' + (dd[1] ? dd  : "0" +dd[0]) ;

const dateStr1 = today;
const date1 = new Date(dateStr1);
const timestamp = date1.getTime();
const inSeconds = Math.floor(timestamp / 1000);

// 雜湊的資料
const dataToHash = 'turingAPI' + inSeconds;
let hash = CryptoJS.SHA256(dataToHash);
const turing_token = hash.toString(CryptoJS.enc.Hex);
console.log("today: " + today)
console.log("timestamp: " + timestamp)
console.log("inSeconds: " + inSeconds)
console.log("dataToHash: " + dataToHash)
console.log("turing_token: " + turing_token)
