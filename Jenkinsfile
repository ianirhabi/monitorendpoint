pipeline {
   agent none

   stages {
      stage('Build') {
        steps {
          echo 'Building...'
          echo "Running ${env.BUILD_ID} ${env.BUILD_DISPLAY_NAME} on ${env.NODE_NAME} and JOB ${env.JOB_NAME}"
        }
   }
   stage('Test') {
     steps {
        echo 'Testing...'
     }
   }
   stage('Deploy') {
     steps {
       echo 'Deploying...'
       bash 'go version'
     }
   }
  }
}

// pipeline {
//     agent { docker 'docker:stable' } 
//     stages {
//         stage('Example Build') {
//             steps {
//                 sh 'ls'
//                 sh 'docker ps'
//             }
//         }
//     }
// }