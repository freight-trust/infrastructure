cloud instance name tool
^ec2-(ue1|uw1|uw2|ew1|ec1|an1|an2|as1|as2|se1)-([1-2]{1})([a-c]{1})-(d|t|s|p)-([a-z0-9\-]+)$


^arn:(?P<IEAn>[^:\n]*):(?P<Service>[^:\n]*):(?P<Region>[^:\n]*):(?P<AccountID>[^:\n]*):(?P<Ignore>(?P<ResourceType>[^:\/\n]*)[:\/])?(?P<Resource>.*)$

# ARN Format Templates
arn:partition:service:region:account-id:resource
arn:partition:service:region:account-id:resourcetype:resource
arn:partition:service:region:account-id:resourcetype/resource
# Amazon Simple Storage Service (S3)
arn:aws:s3:::my_corporate_bucket/Development/*
