# CI/CD with Golang

[![Go](https://github.com/devenes/containerization-dummy-go/actions/workflows/go.yml/badge.svg)](https://github.com/devenes/containerization-dummy-go/actions/workflows/go.yml)

## Part 1

- Since we want to trigger the CI/CD process with the push we will perform within the GitHub ecosystem, we have pushed the source code to our own repo.

- We used GitHub Actions, the Pipeline as Code tool offered by GitHub, to trigger the process with the push commit.

- A script was written to perform the Docker build process with the trigger and to give the version number.

```
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .

RUN go build -o /dummy-app
EXPOSE 8080

CMD ["/dummy-app"]
```

- The username and password of the Docker Hub Registry, which is the environment we will push our Docker images to after we build, are defined as secrets in GitHub Repo Settings and used as environment variables in the GitHub Actions pipeline.

- To trigger our Jenkins pipeline, we add the curl command to the end of our GitHub Actions pipeline and our token that we defined in Jenkins to send a request to the webhook link of the server.

  - Note: Token is defined as environment variable in secrets from GitHub repo settings.

In our Jenkins pipeline, we first stop and delete the running container and clean up the old images. Then we download the Image that we have determined with the current tag number and run it as a container.

We give 80 as the port number in the docker run command so that our Jenkins running on 8080 and our Go application do not conflict with each other.

The output of our operations looks like this:

Our image that we pushed to Docker Hub after we got the Build:

## Part 2

Before starting to operate on CodeBuild, we give some permissions to reach the ECR and to examine the post-process logs in detail:

We selected our repo by logging into CodeBuild with GitHub. Then we edit the build spec file with the links we got from the ECR settings:

```
version: 0.2

phases:
  pre_build:
    commands:
      - echo Logging in to Amazon ECR...
      - aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com
  build:
    commands:
      - echo Build started on `date`
      - echo Building the Docker image...
      - docker build -t mydockerrepo .
      - docker tag mydockerrepo:latest 32976713052.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$IMAGE_REPO_NAME:$IMAGE_TAG
  post_build:
    commands:
      - echo Build completed on `date`
      - echo Pushing the Docker image...
      - docker push 32976713052.dkr.ecr.us-east-1.amazonaws.com/mydockerrepo:latest
```

We see that our action steps have been successfully completed:

![image](</readme-images/1%20(6).png>)

At the end of our Image Build process, we come to the repo we opened on ECR to see if our push to ECR was successful:

![image](</readme-images/1%20(7).png>)

<!-- # CI/CD with Golang

## 1. Kısım

- GitHub ekosistemi içinde gerçekleştireceğimiz Push işlemiyle birlikte CI/CD sürecini tetiklemek istediğimiz için kaynak kodu kendi repomuza pushladık.

- Push commiti ile birlikte süreci tetikleme işlemini başlatmak için GitHub’ın sunduğu Pipeline as Code aracı olan GitHub Actions’ı kullandık.

- Docker build işleminin tetikleme ile birlikte gerçekleşmesi ve versiyon numarasının verilmesi için script yazıldı.

- Docker imajlarımızı build aldıktan sonra push’layacağımız ortam olan Docker Hub Registry’nin kullanıcı adı ve şifresi GitHub Repo Ayarları içinde secrets olarak tanımlandı ve GitHub Actions pipeline’ı içinde ortam değişkeni olarak kullanıldı.

- Jenkins pipeline’ımızı tetiklemek için server’a ait webhook linkine istek göndermek amacıyla GitHub Actions pipeline’ımızın sonuna curl komutu ve Jenkins’e de tanımladığımız token’ımızı ekliyoruz.

  - Not: Token GitHub repo ayarlarından secretlar içerisinde ortam değişkeni olarak tanımlandı.

Jenkins pipeline’ımızda öncelikle çalışan konteynırı durdurup siliyoruz ve eski imajları temizliyoruz. Ardından güncel tag numarası ile belirlediğimiz İmajı indirip konteynır olarak çalıştırıyoruz.

Docker run komutu içinde port numarası olarak 80’i veriyoruz ki 8080’de çalışan Jenkins ile Go uygulamamız birbiriyle çakışmasın.

İşlemlerimizin çıktısı şu şekilde karşımıza geliyor:

Build aldıktan sonra Docker Hub’a pushladığımız imajımız:

## 2. Kısım

CodeBuild üzerinde işlem yapmaya başlamadan önce ECR’a ulaşması için ve işlem sonrası logların detaylı olarak incelenebilmesi için bazı yetkiler veriyoruz:

CodeBuild üzerinde GitHub ile oturum açarak repomuzu seçtik. Sonrasında build spec dosyasını ECR ayarlarından aldığımız linklerle düzenliyoruz:

İşlem adımlarımızın başarıyla tamamlandığını görüyoruz:

Image Build işlemimiz sonunda ECR’a push aktivitemiz başarılı olmuş mu, bunu görmek için ECR üzerinde açtığımız repoya geliyoruz: -->
