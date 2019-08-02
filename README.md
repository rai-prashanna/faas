# SIMPLE Function as a service (FaaS) IMPLEMENTATION
* [Function as a service](https://en.wikipedia.org/wiki/Function_as_a_service) 


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 


### Prerequisites
What things you need to install the software and how to install them

* [install GOLANG](https://golang.org/doc/install) 
* [govendors - To manage dependecies](https://github.com/kardianos/govendor) 
* [install Docker](https://docs.docker.com/install/)
* [Post-installation steps for Linux](https://docs.docker.com/install/linux/linux-postinstall/)
* [Install Docker Compose](https://docs.docker.com/compose/install/)


### To run in your local machine

```
$ git clone https://github.com/rai-prashanna/faas/
$ cd <Working Directory>/faas/
$ docker-compose up
$ use this url on web browser (http://localhost:8080/factorial?num=3 or http://localhost:8080/dig?url=www.wwe.com)
```
[http://localhost:8080/factorial?num=3](http://localhost:8080/factorial?num=3)
<br />
[http://localhost:8080/dig?url=www.wwe.com](http://localhost:8080/dig?url=www.wwe.com)




### little demo

![alt text](https://github.com/rai-prashanna/faas/blob/master/success.png)
<br />
![alt text](https://github.com/rai-prashanna/faas/blob/master/output1.png)
<br />
![alt text](https://github.com/rai-prashanna/faas/blob/master/output2.png)




## Authors

* **Patrick Rai** 


## License

[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/rai-prashanna/faas/blob/master/LICENSE)
This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/rai-prashanna/faas/blob/master/LICENSE) file for details

## Acknowledgments

* Thanks to Modular Finance for assignment

