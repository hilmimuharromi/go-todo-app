def agentLabel
if (BRANCH_NAME == "staging") {
    agentLabel = "agent-staging"
} else {
    agentLabel = "agent-production"
}


pipeline {
    agent {label agentLabel}
    stages {
        stage('Pulling latest code') {
            steps {
                echo " ============================== Starting Pulling latest Code  $agentLabel=============================="
            checkout scmGit(branches: [[name: BRANCH_NAME]], extensions: [], userRemoteConfigs: [[credentialsId: 'hilmi-github', url: 'https://github.com/hilmimuharromi/go-todo-app']])
            }
        }
          stage('Pre build') {
            steps {
                echo " ============================== Starting Pre Build =============================="
                sh 'docker stop container-go-todo-app'
                sh 'docker rmi go-todo-app:latest'
            }
        }
        stage('Build') {
            steps {
                echo " ============================== Starting Build =============================="
            sh 'docker build -t go-todo-app:latest .'
            }
        }
        stage('Running Container') {
            steps {
                echo " ============================== Starting Run Project =============================="
            sh 'docker container run -d --name container-go-todo-app --network host --rm --tty -e PORT=8000 -p 8000:8000 go-todo-app:latest'
            }
        }
    }
}
