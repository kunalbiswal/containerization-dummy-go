pipeline {
  agent any
  stages {
    stage("Stop Running Container") {
      steps {
        sh "docker stop dummy-go"
      }
    }
    stage("Delete Container") {
      steps {
        sh "docker rm dummy-go"
      }
    }
    stage("Delete Image") {
      steps {
        sh "echo y | docker image prune -a"        
      }
    }
    stage("Pull Image") {
      steps {
        sh """
          docker pull devenes/dummy-go:$version
        """
      }
    }
    stage("Run Container") {
      steps {
        sh """
          docker run --name dummy-go -d -p 80:8080 devenes/dummy-go:$version
        """
      }
    }    
  }
}
