# leetcode

### Memoization Recipe
1. Make it work.
• visualize the problem as a tree
• implement the tree using recursion
• test it
2. Make it efficient.
• add a memo object
• add a base case to return memo values
• store return values into the memo

### Tabulation Recipe
• visualize the problem as a table
• size the table based on the inputs
• initialize the table with default values
• seed the trivial answer into the table
• iterate through the table
• fill further positions based on the current position

## Tips to Consider

```
If input array is sorted then
    - Binary search
    - Two pointers

If asked for all permutations/subsets then
    - Backtracking

If given a tree then
    - DFS
    - BFS

If given a graph then
    - DFS
    - BFS

If given a linked list then
    - Two pointers

If recursion is banned then
    - Stack

If must solve in-place then
    - Swap corresponding values
    - Store one or more different values in the same pointer

If asked for maximum/minimum subarray/subset/options then
    - Dynamic programming

If asked for top/least K items then
    - Heap

If asked for common strings then
    - Map
    - Trie

Else
    - Map/Set for O(1) time & O(n) space
    - Sort input for O(nlogn) time and O(1) space
```

## References

<https://github.com/seanprashad/leetcode-patterns>

<https://github.com/PinkyJie/leetcode-patterns>

<https://neetcode.io/>

-----------

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

------------------------------------------
From call centre agent and not knowing how to code to Software Engineer at Microsoft 
Here is how: 

8 years ago as a call center agent, Goran was working to resolve people’s issues with their phones. He would go to school early in the morning, then off to work right after for an 8 hour shift and taking hundreds of phone calls each day. 

At 12 am he would go home exhausted, but deep down he had a dream to follow. He wanted to be among the best software engineers in the world. So many people doubted him. Why wouldn’t they, he didn’t know how to code at all at this time. 

He was always hungry to learn to reach his dreams. He looked at challenges as something to overcome not as something to make him give up.

He sought guidance and support both off and online. He surrounded himself with highly motivated and smart people.

He would spend weekends and summer holidays teaching himself the basics of coding. Then to different programming languages. 

He once worked for company with a very low paying job that was just covering food expenses, just because he wanted to sit beside their developers and learn how they build websites. 

His main aim was to learn and grow his career, making money wasn't a priority as long as he was making enough to cover his expenses. 

He has put in thousands of hours to reach where he is and knowing him closely I am certain only the sky is the limit.

Most importantly he reached here not by only focusing on his career. He is a great husband and father, and I know from experience a great friend and neighbour as well.

His most important advise to people who are starting their career journey are: 
 
 - Have big dreams
 - Keep learning and always stay curious
 - Be patient
 - Have faith in yourself  yourself especially when next steps are not clear
 - Early in the career choose experience over money

Always have big dreams and work towards them accordingly. People around you may not understand.

### Resources
https://github.com/jwasham/coding-interview-university
https://github.com/ashishps1/awesome-system-design-resources
https://github.com/ashishps1/awesome-low-level-design
https://github.com/ashishps1/awesome-coding-projects
https://github.com/ashishps1/awesome-leetcode-resources
https://github.com/ashishps1/awesome-behavioral-interviews
https://leetcode.com/discuss/study-guide/458695/Dynamic-Programming-Patterns
https://leetcode.com/discuss/study-guide/1688903/Solved-all-two-pointers-problems-in-100-days
https://leetcode.com/discuss/study-guide/786126/Python-Powerful-Ultimate-Binary-Search-Template.-Solved-many-problems
https://leetcode.com/discuss/study-guide/655708/Graph-For-Beginners-Problems-or-Pattern-or-Sample-Solutions
