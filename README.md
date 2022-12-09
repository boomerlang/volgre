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
