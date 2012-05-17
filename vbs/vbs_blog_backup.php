<?php

$pdo = new pdo("sqlite:data/demon.tw");
//$pdo->exec("create table vbs(id integer primary key autoincrement,title varchar(255),body text,url unique)");

function fixbody($body){
$body = trim($body);

return str_replace(array("作者:","版权:","协议条款。"),
array("\n作者:","\n版权:","协议条款。\n\n"),$body);
}



for($page=39;$page>0;$page--){
  try{
   get($page);  
  } catch(exception $e){
     
  }
}
function get($page) {
 global $pdo;
 
$url = "http://demon.tw/page/$page?s=vb";

$doc = new domdocument();
@$doc->loadhtmlfile($url);
$x= new domxpath($doc);

$links = $x->query("//div[@id='content']//h3/a");

foreach($links as $link){
  $title = $link->textContent;
  
 $page_url = $link->getAttribute("href");
 $page_url =trim($page_url ); 
  if(!$page_url){ continue;}
  
  $doc2 = new domdocument();
  @$doc2->loadhtmlfile($page_url );
  $x2 = new domxpath($doc2);
  $ps = $x2->query("//div[@class='entry']");
  foreach($ps  as $p){
      $body = $p->textContent;         
  }
  $body = fixbody($body);
  
  $s=$pdo->prepare("insert into vbs(title,body,url) values(:title,:body,:url)");
  $s->execute(array('title'=>$title,'body'=>$body,'url'=>$page_url));
  
}

}