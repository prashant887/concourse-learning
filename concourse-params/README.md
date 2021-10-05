# concourse-params

### Set Params with crendtial file 
fly -t tutorial set-pipeline -p con-params-pipeline -c pipeline-parameters.yml -l credentials.yml 

### Set Params Command Line 

fly -t tutorial set-pipeline -p con-params-pipeline -c pipeline-parameters.yml -v username=admin -v password=adminpass
### Trigger Job 
fly -t tutorial trigger-job -j  con-params-pipeline/show-crendetials-job -w 
