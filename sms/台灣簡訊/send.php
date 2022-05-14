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
$account = '';
// password
$password = '';
// receiver
$tel = '';
// content
$content = urlencode('This is 台灣簡訊平台寄送測試, 1!@#$%^&*()_=//\/');
$url = "https://api.twsms.com/json/sms_send.php?username=$account&password=$password&mobile=$tel&message=$content";


var_dump($url);

$opts = array('https' =>
    array(
        'method' => 'GET',
        'header' => 'Content-type: application/x-www-form-urlencoded',
    )
);
$context = stream_context_create($opts);
$result = file_get_contents($url, false, $context);
var_dump($result);

