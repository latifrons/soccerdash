# SoccerDash

SoccerDash is the aggregate dashboard for demonstrating data from multiple peers/nodes on one web page.

It receives messages from multiple nodes through a simple socket server, and shows them on web page through websocket.

SoccerDash is extremely helpful and easy to use during multiple nodes debugging, such as cryptocurrency nodes
sync debugging.  

## Start a log aggregating server
```
python3 multiple_logger.py
```
It will starts an aggregate webserver to collect logs from different nodes on port 1088 and a websocket for webpage on 1099

## Start a metric aggregating server

```
python 3 multiple_metric.py
```

It will starts an aggregate webserver to collect metrics from different nodes on port 1088 and a websocket for webpage on 1099

## Embed into client

```
Some usage TODO
```

## Authors

* **latifrons**


## License

This project is licensed under the MIT License

## Acknowledgments

