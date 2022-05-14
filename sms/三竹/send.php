<?php

/** 
 * PHP version 7.4.28
 * 三竹簡訊平台寄送測試

 * @category Components
 * @package  SmsMessage
 * @author   Chiehting-Lee <ting911111@gmail.com>
 * @license  https://www.apache.org/licenses/LICENSE-2.0 APACHE LICENSE, VERSION 2.0
 * @link     https://chiehting.com
 **/

// account
$id = '';
// password
$password = '';
// receiver
$tel = '';
// content
$content = urlencode('This is 三竹簡訊平台寄送測試, 1!@#$%^&*()_=//\/');
$url = "https://sms.mitake.com.tw/b2c/mtk/SmSend?";
$url .= 'CharsetURL=UTF-8';

// parameters
$data = "username=$id";
$data .= "&password=$password";
$data .= "&dstaddr=$tel";
$data .= "&smbody=$content";

var_dump($url);
var_dump($data);

$curl = curl_init();
// 設定curl網址
curl_setopt($curl, CURLOPT_URL, $url);
// 設定Header
curl_setopt($curl, CURLOPT_HTTPHEADER, array("Content-type: application/x-www-form-urlencoded"));
curl_setopt($curl, CURLOPT_POST, true);
curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
curl_setopt($curl, CURLOPT_HEADER,true);
// 執行
$output = curl_exec($curl);
curl_close($curl);
echo $output;
