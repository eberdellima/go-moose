pipeline {
  agent { docker { image: 'node:14-apline' } }
  stages {
    stage('build') {
      steps {
        timeout(time: 2, unit: 'MINUTES') {
          echo 'Hello'
        }
      }
    }
  }
  post {
    always {
      echo 'piline finished ...'
    }
    success {
      echo 'pipeline finished successfully'
    }
    failure {
      echo 'pipeline failed'
    }
  }
}