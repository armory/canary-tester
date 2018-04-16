node {
    checkout scm

    stage("Build deb") {
        sh("scripts/build-deb.sh")
    }

    if (env.BRANCH_NAME == "master") {
        stage("Publish Deb Package") {
            sh '''
              . /mnt/secrets/bintray/bintray
              scripts/publish-deb.sh
            '''
        }
    }

    stage("Archive Artifacts") {
           archiveArtifacts artifacts: 'build/distributions/*.deb', fingerprint: true
    }

}