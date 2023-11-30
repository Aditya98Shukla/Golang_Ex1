Write a program to read the system details ("lscpu" command), disk space, and top 10 process details (which is consuming high resource , for example use "top" command) and construct the json with all the details and expose it through endpoint by using http or any other as per your convinent.

exaple O/p:
curl http://localhost:8080/systemstats - if you do the curl command on this endpoint it should provide output.
{
"Architecture": "x84_64",
"CPU" : "4",
...
...
...
"process": [{
    "Pid":
    "cpu":
    "usr":
    ...
    ... (add all the fields from the "top" command)
    },
    {
    "Pid":
    "cpu":
    "usr":
    }],
"disk":{

    }
"totalDiskSpace": "100GB",
"usedSpace": "50GB",
"freeSpace": "50GB",
}


Add as many field to the struct related to the system details.
