pipeline {
   agent any

    stages {
        stage('Build') {
            if(env.BRANCH_NAME == 'master2'){
                steps {
                echo 'Building...'
                echo "Running ${env.BUILD_ID} ${env.BUILD_DISPLAY_NAME} on ${env.NODE_NAME} and JOB ${env.JOB_NAME}"
                }
            }
        }
        stage('Test') {
            if(env.BRANCH_NAME == 'master2'){
                steps {
                    echo 'Testing...'
                }
            }
        }
        stage('Deploy') {
            steps {
            echo 'Deploying...'
            }
        }
    }
}