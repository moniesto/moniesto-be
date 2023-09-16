package util

const JOB_TYPE_EVERY_SECOND = "* * * ? * *" // does not work because of the version update of cron library (github.com/robfig/cron -> github.com/robfig/cron/v3)
const JOB_TYPE_EVERY_MINUTE = "* * * * *"
const JOB_TYPE_EVERY_5TH_MINUTE_OF_HOUR = "5 * * * *"
const JOB_TYPE_EVERY_30TH_MINUTE = "*/30 * * * *"
const JOB_TYPE_EVERY_HOUR = "0 * * * *"
const JOB_TYPE_EVERY_12AM = "0 0 * * *"
const JOB_TYPE_EVERY_1AM = "0 1 * * *"
