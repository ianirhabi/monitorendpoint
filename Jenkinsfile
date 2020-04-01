// pipeline {
//    agent none

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
//        go version
//      }
//    }
//   }
// }

pipeline {
    agent { docker 'maven:3-alpine' } 
    stages {
        stage('Example Build') {
            steps {
                sh 'mvn -B clean verify'
            }
        }
    }
}