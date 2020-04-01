pipeline {
   agent any

//    stages {
//       stage('Build') {
//         steps {
//           echo 'Building...'
//           echo "Running ${env.BUILD_ID} ${env.BUILD_DISPLAY_NAME} on ${env.NODE_NAME} and JOB ${env.JOB_NAME}"
//         }
//    }
//    stage('Test') {
//      steps {
//         echo 'Testing...'
//      }
//    }
//    stage('Deploy') {
//      steps {
//        echo 'Deploying...'
//      }
//    }
//   }
  // Run on an agent where we want to use Go
    node {
        // Install the desired Go version
        def root = tool name: 'Go 1.8', type: 'go'

        // Export environment variables pointing to the directory where Go was installed
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
        }
    }
}