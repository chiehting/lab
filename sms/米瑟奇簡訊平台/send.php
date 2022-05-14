<?php

/** 
 * PHP version 7.4.28
 * Message 平台簡訊寄送測試 

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
$content = urlencode('This is 米瑟奇平台簡訊測試');
$url = "http://api.message.net.tw/send.php?id=$id&password=$password&tel=$tel&msg=$content&encoding=utf8";
$sendDate = '';

echo mb_detect_encoding($content);

if (empty($sendDate)) {
    $url = $url . '&mtype=G';
} else {
    $url = $url . '&sdate=' . $sendDate;
}

var_dump($url);

$curl = curl_init();
// 設定curl網址
curl_setopt($curl, CURLOPT_URL, $url);
// 設定Header
curl_setopt($curl, CURLOPT_HTTPHEADER, array("Content-type: application/x-www-form-urlencoded"));
curl_setopt($curl, CURLOPT_POST, false);
curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
curl_setopt($curl, CURLOPT_HEADER,true);
// 執行
$output = curl_exec($curl);
curl_close($curl);
echo $output;

