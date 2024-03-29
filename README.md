# SIMPLE Function as a service (FaaS) IMPLEMENTATION
* [Function as a service](https://en.wikipedia.org/wiki/Function_as_a_service) 

# Background
* So this docker thing, as well as ​ function as a service, FaaS, ​ seems to getting some traction. It would be
  nice if we could have our own ​ FaaS ​ infrastructure in place instead of paying Amazon for it. There are
  some alternatives out there e.g. OpenFaaS. But we have a bad case of “Not Invented Here” syndrome,
  so we’d rather build it ourselves.


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
$ docker-compose build --no-cache
$ docker-compose up
$ use this url on web browser (http://localhost:8080/function/:factorialservice?num=3 or http://localhost:8080/function/:digservice?url=www.wwe.com)
```

[http://localhost:8080/function/:factorialservice?num=3](http://localhost:8080/function/:factorialservice?num=3)
<br />
[http://localhost:8080/function/:digservice?url=www.wwe.com](http://localhost:8080/function/:digservice?url=www.wwe.com)

### Note

```
$ In first attempt of docker-compose up command,
$ the process might be very slow. Since it downloads all dependencies from remote.
$ So be patient to let docker-compose build images and docker-compose to run all services
```


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

