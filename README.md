# Integra Hackathon

2025-04-03

This is a fun task to challenge your golang Kungfu. Submit your implementation by 1.45 pm on 04/03/2025 and become part of the leaderboard!

The text file contains temperature values for a range of weather stations. Each row is one measurement in the format <string: station name>;<double: measurement>, with the measurement value having exactly one fractional digit. The following shows ten rows as an example:

```data
Hamburg;12.0
Bulawayo;8.9
Palembang;38.8
St. John's;15.2
Cracow;12.6
Bridgetown;26.9
Istanbul;6.2
Roseau;34.4
Conakry;31.2
Istanbul;23.0
```

The task is to write a golang program which reads the file, calculates the min, mean, and max temperature value per weather station, and emits the results on stdout like this (i.e. sorted alphabetically by station name, and the result values per station in the format <min>/<mean>/<max>, rounded to one fractional digit):

```output
{Abha=-23.0/18.0/59.2, Abidjan=-16.2/26.0/67.3, Abéché=-10.0/29.4/69.0, Accra=-10.1/26.4/66.4, Addis Ababa=-23.7/16.0/67.0, Adelaide=-27.8/17.3/58.5, ...}
```

The file contains 1 billion rows.  In order to facilitate the testing , you are provided with sample files with 100k records and 10 million records. The actual file with 1 billion records is on the Test VM, since this is a 15 gb file, we don’t recommend you download this to your machine.  Download the test files from the below link.

https://integrapartners.sharepoint.com/sites/HackathonDocs/Shared%20Documents/Forms/AllItems.aspx

The specs of the Test VM  is a d32 v6(32 vcpus,128gb ram)

You will have to submit the golang source to your personal git hub repository and share the link with Shiva. He will run it on the VM and publish the results.

The folder structure of the code must be structured as shown below

```
pkg
    ├── cmd
    │   └── main.go
    └── data
        └── measurements.txt
```

 