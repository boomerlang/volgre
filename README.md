Multi Rule Engine Processor
----------------------------
----------------------------


Install volgre application + dependencies:
------------------------------------------


$ git clone https://github.com/boomerlang/volgre


Run the application:
--------------------


$ cd volgre


$ go build


$ ./volgre  -host _your server ip address_ -port 8082 &


Test the server
---------------


In other console:


$ cd volgre


$ curl -v -X POST http://_your server ip address_:8082/run/engine/credit_card --data @input.json


Output Data
-----------

Uses 'volgre' namespace in the response json data that appends to the original json.


On an Intel(R) Core(TM) i5-5675R CPU @ 3.10GHz, 8GB RAM rule engine execution time is few hundred microseconds.


On other configuration the timings might vary.

Quick Test
----------

$ curl -X POST http://141.147.22.26:8084/run/engine/credit_card --data @input.json


Endpoints
---------

1. Runs a rule engine with the corresponding loaded rule set:

/run/engine/{engine_name}


2. Refreshes a rule engine with a new set of rules

/refresh/engine/{engine_name}


3. Show the version of the current rule set of a given rule engine

/version/engine/{engine_name}


Measurements
------------


<table>

<tr>
<td>No. of rules</td><td>Engine execution time (ms)</td><td>Total execution time  (ms)</td><td>Refresh time</td>
</tr>

<tr>
<td>5</td><td>0.6</td><td>4</td><td>5</td>
</tr>

<tr>
<td>10</td><td>         0.8   </td><td>              5</td><td>7</td>
</tr>
<tr>
<td>50</td><td>         2.5   </td><td>              8</td><td> 12</td>
</tr>

<tr>
<td>100</td><td>           5   </td><td>             10</td><td> 20</td>
</tr>

<tr>
<td>300</td><td>          17   </td><td>             20</td><td> 40</td>
</tr>

<tr>
<td>600</td><td>          40   </td><td>             45</td><td> 90</td>
</tr>
<tr>
<td>1000</td><td>          90   </td><td>             95</td><td>145</td>
</tr>
<tr>
<td>1500</td><td>         150   </td><td>            148</td><td>205</td>
</tr>
<tr>
<td>2000</td><td>         255   </td><td>            250</td><td>269</td>
</tr>
<tr>
<td>3000</td><td>         495   </td><td>            485</td><td>403</td>
</tr>
</table>

Examples
--------

$ curl -X POST http://141.147.22.26:8084/run/engine/credit_card --data @input.json


$ curl -X GET http://141.147.22.26:8084/refresh/engine/credit_card --data @input.json


$ curl -X GET http://141.147.22.26:8084/version/engine/credit_card --data @input.json

Misc
------

To generate a set of 1200 distinct rules use the following shell script:

$ bash gen_rules.sh 200 > path_to/rule_file.grl


