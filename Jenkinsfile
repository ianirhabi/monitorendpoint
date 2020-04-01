pipeline {
   agent any

    stages {
        stage ('Speak') {
            when {
                expression { env.NODE_NAME == 'master2' }
                    steps {
                        echo "Hello, bitwiseman!"
                    }
            }
        }
    }
}