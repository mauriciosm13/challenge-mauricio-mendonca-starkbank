Stark Bank Challenges
Welcome to the Stark Bank Challenges repository!
This project involves two tasks that simulate real-world financial operations using the sandbox environment.


Challenge 1: Automated Invoicing
Overview
Objective: Issue invoices to random individuals every 3 hours for a duration of 24 hours.
Environment: All operations will be conducted within the Stark Bank Sandbox environment, which will automatically process some of the invoices.
Expected Outcome:
Logs showing the invoices sent, including timestamps, recipient details, and amounts.
Some invoices automatically paid by the sandbox environment

Challenge 2: Webhook Handling and Transfer
Overview
Objective: Upon receiving an invoice payment confirmation via webhook, transfer the net amount (after any fees) to a specified bank account.
API:
Use the Stark Bank API to create and send these invoices.

Steps to Implement:
Webhook Setup:
Configure the server to receive webhooks from Stark Bank for invoice payments. Using NGROK.
Transfer Execution:
Use the Stark Bank API to initiate a transfer
Expected Outcome:
A system that automatically receivies invoice payments and performs corresponding bank transfers, with detailed logs for auditing.
