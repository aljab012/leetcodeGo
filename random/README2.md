A Netflix engineer did 7 onsites interviews for 7 tech companies.  Received 7 offers ranging $325k - $800k a year. He/she shared the resources in Blind that helped her/him to be successful. Here they are:

1. Leetcode
Always ask clarifying questions, they are meant to be vague.

Tech Interview Handbook (solve 50 suggested LCs multiple times): https://lnkd.in/dbK9TasP

Grokking Coding Interview:
https://lnkd.in/d4YzR_3d

Solve each question yourself before looking at answers. After struggling with certain topics, found some youtube resources that were really helpful.

- Graph problems: https://lnkd.in/dT4brkwG
- Kahn’s topsort: https://lnkd.in/dc8VzPx7
- Dijkstra’s shortest path: https://lnkd.in/djWD4ncp
- Kruskal’s MST: https://lnkd.in/dUaF7dDr
- Union find: https://lnkd.in/dmZSyGXq
- Dynamic programming: https://lnkd.in/d85BHsQp

I would say 1/2 of my coding questions was graph related, and I used topsort, MST, union find during my interviews so its worth knowing them well.

Great resource on Binary Search (esp. for harder variants): https://lnkd.in/dPVikn9e

Quite a few LC hards involve binary search as part of its solution, but its non-obvious.

At staff level, you are expected to have great communication, come up with one or more high level solutions in reasonable time with space and time complexity, implement mostly correct, and check for errors independently.

2. SYSTEM DESIGN

Always ask clarifying questions, they are meant to be vague.

No. 1 tip: Pay for mock interviews, take extensive notes, study, pay for more interviews. I used igotanoffer.com. Try booking different interviewers, they have different expectations, communication, friendliness.

Grokking system design: https://lnkd.in/dKBtGDDG

It’s good to complete grokking, but in practice its a little too shallow for staff level.

Watch all videos and take notes from this channel: https://lnkd.in/d32sycXJ

Most other channels like Exponent are actually not great examples of a good interviewee for staff level.

Read papers on Dynamo, BigTable, MapReduce, Cassandra, Raft, Lamport Clocks, etc.

A lot of material written by this Medium writer is great: https://lnkd.in/d_4yxAA7

Memorize Jeff Dean’s latency numbers: https://lnkd.in/dq42yWsa

If you have time, read DDIA: https://dataintensive.net

I memorized mnemonics for structure of how I execute my interview, as well as for each topic. Like OSI model, caching strategies, load balancing strategies, database types, etc. You need to make up your own in order to ingrain them.


Remaining tips are in comments:

3. BEHAVIORAL

Prepare an introduction about yourself. Keep it short and sweet, highlighting background and - a major accomplishment. (Don’t be afraid to sell yourself)

Grokking the behavioral: https://www.educative.io/courses/grokking-the-behavioral-interview

The resource itself isn’t that useful but should prepare answers from your own experience and recording yourself. Take time to reflect on your past few years and make sure you have answers about conflict, projects you led, giving/receiving feedback, and challenges you solved. Don’t lie.

STAR is a good structure for your stories but I learned a better story shape from this YouTuber.

U shape story: https://youtu.be/hU6BVxtGd5g

Read “Staff Engineer”: https://staffeng.com/book

For each company, make sure you know their products, their company motto and values. I kept notes which I read a few mins before my onsite starts. Prepare plenty of questions to ask.

Number 2 tip: Every system design question has a point of contention. Is it storage? Latency? Where will the system start to break? This is pivotal to figure out asap.

Basic structure:
Requirements (functional / non-functional)
Estimations (focus on contention point)
System API
Data model
High level design (core components)
Detailed design (ask where to dive in)
Bottlenecks (scale, redundancy, SPOFs, metrics, logs, alerts, dashboards, pagerduty, deployment, failure scenarios)

As staff level interviewee, you are expected to lead the interview and provide alternatives, weigh pros & cons of each approach and commit to one with justification. Outside of typical system design topics, staff level also needs to consider costs, system complexity, and business aspects (e.g. CDN too expensive for free tier).

For most interviewers, you should be familiar with at least one database technology for each category (key value, document, columnar, relational, etc). For Google, they typically want you to design from fundamentals instead of using off the shelf technologies.

Number 3 tip: Listen intently to interviewer when they speak and take their hints.
