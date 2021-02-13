pipeline {
  agent { docker { image: 'node:14-apline' } }
  stages {
    stage('build') {
      steps {
        sh 'npm --version'
      }
    }
  }
}