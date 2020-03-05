{
    "parameters": {
        "commit": $COMMIT,
        "environment": "lee",
        "success_message": "Successful deployment",
        "migrate": {
            "migrate": false,
            "task_def": $MIGRATION_TASKDEF_ARN
        },
        "apps": [
            {
                "name": "app-lee",
                "service_deploy": true,
                "task_def": $APP_TASKDEF_ARN,
                "deployment_group": "app-lee-deployment-group"
            }
        ]
    }
}