import com.sap.piper.BuildTool
import com.sap.piper.DownloadCacheUtils
import groovy.transform.Field

import static com.sap.piper.Prerequisites.checkScript

@Field String STEP_NAME = getClass().getName()
@Field String METADATA_FILE = 'metadata/whitesource.yaml'

//Metadata maintained in file project://resources/metadata/mavenExecuteIntegration.yaml

void call(Map parameters = [:]) {
    final script = checkScript(this, parameters) ?: this
    parameters = DownloadCacheUtils.injectDownloadCacheInParameters(script, parameters, BuildTool.MAVEN)

    List credentials = [
        [type: 'token', id: 'orgAdminUserTokenCredentialsId', env: ['PIPER_orgToken']],
        [type: 'token', id: 'userTokenCredentialsId', env: ['PIPER_userToken']],
    ]
    piperExecuteBin(parameters, STEP_NAME, METADATA_FILE, credentials)
}
