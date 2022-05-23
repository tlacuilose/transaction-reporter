# Transaction reporter

Reads a csv file and sends an email with a report.

## Code interface

The user has available a single endpoint that receives an account and the email of the receiver:

POST: /email?account=[account]&email=[receiver-email]

The account is the name of the csv file, example account 001 has file 001.csv.
In ./store there are 2 available files, and only these two can be called with the previous endpoint. This means account has only two possible values: 001 or 002

The receiver-email can be any email.

If the call was successful the email should be sent and a message is returned confirming it was sent. Else a message error and an http error are received.

The project was intended to run on an AWS Lambda, its structure is divided in three layers: data, domain, and presentation.
An API Gateway was configured to accept only the previous endpoint and validate its 2 paramaters, these endpoint triggers the execution of the lambda function.
The 2 available files where manually uploaded to S3, this way the lambda execution can access these files.
SendGrid was used as the email service, AWS Simple Email Service was implemented but it could only run on sandbox mode and configuring new emails was not
ideal. Then, the flow of the project is as follows:

- User calls the endpoint giving and account and an email.
- The API Gateway triggers the lambda function.
- The lambda downloads the account.csv file from S3.
- The lambda generates a summary.
- The summary is sent by the lambda function to the receiver-email using SendGrids API.

Tests are available for most of the available packages. Run tests with:

```bash
go test ./...
```

## How to run

The project has been deployed on a lambda function in AWS, following these steps:

Build the project for linux

```bash
env GOOS=linux GOARCH=amd64 go build -o build/main
```

Zip the executable

```bash
zip -j function.zip build/main
```

Finally, manually upload the zip file using the AWS Console.

The endpoint can be accessed using the following curl command

```bash
curl -X POST -H "x-api-key: PFlbaD9WJs2htnDOBH1HY4XPcBuD9xgI2hT5okRC" https://jqaxh39mac.execute-api.us-west-1.amazonaws.com/default/transaction-reporter/email\?email\=[receiver-email]\&account\=[account]
```


### Running the project locally

Additionally, a local executable called localexe and be used to process the csv files stored in ./store and still uses sendgrid to send the email.

```bash
./localexe
```

The following endpoint can be use to run locally:

```bash
curl -X POST http://localhost:1323/email\?account\=[account]\&email\=[receiver-email]
```

# Result

An email should be sent to the receiver. The email may be in the spam folder.
A message of email sent! will be shown as reponse to the POST request.

This is an example of the content of the email.

<style>
  .parent {
    display: grid;
    place-items: center;
  }

  .card {
    width: clamp(23ch, 50vw, 46ch);
    display: flex;
    flex-direction: column;
    padding: 1rem;
  }

  .visual {
      width: 350;
      width: 100vw;
  }

  td {
    text-align: center;
    border: 3px solid #829ee4;
  }

  table {
    width:  350px;
    border: 3px solid #829ee4;
  }

</style>
 <div class=3D"parent white">
  <div class=3D"card purple">
    <h2>Your report for account #002</h2>
    <div class=3D"visual yellow">
      <img src=3D"https://u27020917.ct.sendgrid.net/ls/click?upn=3DbcmkjxXf=
Ew73EaF7gv9-2BZtkQDVWeDYqtbrNaPqNSTdZqsINfi5-2FrZKYccAmJYtBqPz6SPbWTbcAxMRF=
jmsDf76WRK7wpcBIpqXP8ILeM7lAStg9pmbAT5bRmb7eL-2FDWw4154CAiHh0NrGgom870IbqkZ=
t3kqLeNizBJKtAW2kjk-3D9KL4_6bNnyRD0lL56PBLMeb178kXlYZamxCAvDRxFgDZYZiDa1mgo=
OlXhlAhWrbt8e8pFO0ktLYz-2BYNQVGQH3C210PcFbDF8ULMf86KEXUvb2s08adcFz3x4w5Zc-2=
F3aqd6-2B-2Fd6gXCmV27ZIycXCfRpFXrIe5Ln1hE5vkqgH-2Bg6Pc2gzjnRa0MDu8JEyXlB3fy=
Y4YmfSF4l4cm6SFFNTedtBA7YVuqvuz-2F6FhQmXQGllix8q4-3D width=3D"350">
    </div>
    <p>Total Balance: 152.000000</p>
    <p>Montly Average Credit: 94.000000</p>
    <p>Montly Average Debit: -43.333333</p>
    <h4>Table of monthly transactions.</h4>
    <table>
      <tr>
        <th>Month</th>
        <th>Transactions</th>
      </tr>
  <tr>
<td>February</td>
<td>2</td>
</tr>
<tr>
<td>March</td>
<td>1</td>
</tr>
<tr>
<td>May</td>
<td>1</td>
</tr>
<tr>
<td>June</td>
<td>1</td>
</tr>
<tr>
<td>July</td>
<td>1</td>
</tr>

    </table>
  </div>
</div>
