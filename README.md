# pezdispenser
leasing service for pez resources

[![wercker status](https://app.wercker.com/status/98966ab2a9c4035ef694b4267e43c719/s/master "wercker status")](https://app.wercker.com/project/bykey/98966ab2a9c4035ef694b4267e43c719)

[![GoDoc](https://godoc.org/github.com/pivotal-pez/pezdispenser?status.png)](http://godoc.org/github.com/pivotal-pez/pezdispenser)



## Running tests / build pipeline locally

```

# install the wercker cli
$ curl -L https://install.wercker.com | sh

# make sure a docker host is running
$ boot2docker up && $(boot2docker shellinit)

# run the build pipeline locally, to test your code locally
$ ./testrunner

```


## Running locally for development

```

# install the wercker cli
$ curl -L https://install.wercker.com | sh

#copy the sample env config file
$ cp sample/myenv.sample myenv # fill in your details in the newly created myenv file

#copy the sample wercker local deployment manifest
$ cp sample/wercker_local_deploy.sample wercker_local_deploy.yml # fill in your details in the newly created file


#copy the sample vcap_services definition
$ cp sample/vcap_services_template.json.sample vcap_services_template.json # fill in your details in the newly created file



# make sure a docker host is running
$ boot2docker up && $(boot2docker shellinit)

# run the app locally using wercker magic
$ ./runlocaldeploy myenv

$ echo "open ${DOCKER_HOST} in your browser to view this app locally"

```
