<?php

function timetaken($fn) {
	$start = microtime(true);
	$fn();
	$stop = microtime(true);
	$del = $stop - $start;
	
	echo("time taken $del\n");
}

define('ITERATIONS', 1000000);

$names = array('a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l');
$array = array();
$total = 0;
foreach($names as $name) {
	$array[] = array('name' => $name, 'weight' => 1000);
	$total += 1000;
}

timetaken(function() {
	global $array;
	for($index = 0; $index < ITERATIONS; $index++) {
		$array[ array_rand($array) ][ 'name' ];
	}	
	echo("simple:   ");
});

function sampleForWeight($index, $array) {
	foreach($array as $item) {
      // less-than so zero-weighted items are ignored
      if($index < $item['weight']) {
          return $item['name'];
      } else {
          $index -= $item['weight'];
      }
  }
}

timetaken(function() {
	global $array, $total;
	for($index = 0; $index < ITERATIONS; $index++)  {
		sampleForWeight(rand(0, $total - 1), $array);
	}
	echo("weighted: ");
});

//
// Output on a mid-2011 MBAir:
//
// simple:   time taken 1.2239990234375
// weighted: time taken 3.7670910358429
//