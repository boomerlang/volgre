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

|---------------------------------------------------------------------------------------------------------|
|No. of rules   |   Engine execution time (ms)   |   Total execution time  (ms)   |     Refresh time (ms) |
|---------------------------------------------------------------------------------------------------------|
|    5          |                0.6             |                     4          |       5               |
|---------------------------------------------------------------------------------------------------------|
|   10          |                0.8             |                     5          |       7               |
|---------------------------------------------------------------------------------------------------------|
|   50          |                2.5             |                     8          |      12               |
|---------------------------------------------------------------------------------------------------------|
|  100          |                  5             |                    10          |      20               |
|---------------------------------------------------------------------------------------------------------|
|  300          |                 17             |                    20          |      40               |
|---------------------------------------------------------------------------------------------------------| 
|  600          |                 40             |                    45          |      90               |
|---------------------------------------------------------------------------------------------------------|
| 1000          |                 90             |                    95          |     145               |
|---------------------------------------------------------------------------------------------------------|
| 1500          |                150             |                   148          |     205               |
|---------------------------------------------------------------------------------------------------------|
| 2000          |                255             |                   250          |     269               |
|---------------------------------------------------------------------------------------------------------|
| 3000          |                495             |                   485          |     403               |
|---------------------------------------------------------------------------------------------------------|


Examples
--------

$ curl -X POST http://141.147.22.26:8084/run/engine/credit_card --data @input.json


$ curl -X GET http://141.147.22.26:8084/refresh/engine/credit_card --data @input.json


$ curl -X POST http://141.147.22.26:8084/version/engine/credit_card --data @input.json

Misc
------

To generate a random number of distinct rules use the following shell script

$ bash gen_rules.sh 200 > path_to/rule_file.grl


