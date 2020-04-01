pipeline {
   agent any

    stages {
        stage ('Speak') {
            when {
                expression { env.NODE_NAME == 'master' }
            }
            steps {
                echo "Hello, bitwiseman!"
            }
        }
    }
}