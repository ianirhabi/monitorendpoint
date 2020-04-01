// pipeline {
//    agent any

//     tools {
//         go 'golang'
//     }

//    stages {
//       stage('Build') {
//         when { branch 'master' }
//         steps {
//           echo 'Building...'
//           echo "Running ${env.BUILD_ID} ${env.BRANCH_NAME} on ${env.NODE_NAME} and JOB ${env.JOB_NAME}"
//         }
//     }

//    stage('Test') {
//      steps {
//         echo 'Testing...'
//      }
//    }

//    stage('Deploy') {
//         steps {
//             withEnv(["PATH+EXTRA=${HOME}/go/bin"]){
//                echo 'Deploying...'
//                sh 'go version'
//             }
//         }
//       }
//   }
// }

pipeline {
    agent { docker 'docker:stable' } 
    stages {
        stage('coba') {
            when { branch 'master' }
            steps { 
                sh 'ls'
                sh 'docker ps'
            }
        }
    }
}


// pipeline {
//    agent any
   
//     tools {
//         go 'golang'
//     }

//     options {
//       disableConcurrentBuilds()
//     }

//    stages {
//       stage('Deploy') {
//         steps {
//             withEnv(["PATH+EXTRA=${HOME}/go/bin"]){
//                echo 'Deploying...'
//                sh 'go version'
//             }
//         }
//       }
//   }
// }