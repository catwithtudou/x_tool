## X(twitter) Tool(Toy)

Features currently ready to be supported:

- [x] Automatic character cutting(Support Chinese characters)
- [ ] End point interaction optimization
- [ ] Automated Tweets
- [ ] optimized typography


## Automatic character cutting

> also support Chinese characters

examples:

1. Execute the main
2. Input the content of the tweet according to the end point prompt
3. Finally, if the content of the tweet exceeds the word limit of the tweet, the program will automatically cut out multiple tweets and return them



```shell
./main
===================================================
Please enter the content of the tweet (finally enter "exit" to end the input):
这会是一条测试文案：以下语句摘自《正见》
潜意识中，我们期待自己会到达不再需要修理任何东西的境界。总有一天，我们会「从此过着快乐的生活」。我们深信「解决」的概念。好像我们所有经历的一切，到这一刻为止的生命，都只是在彩排。盛大的演出还没开始。对大多数的人来说，这种永无休止的处理和重新安排以及更新版本，就是生活的定义。事实上，我们是在等待生命开始。
我们往往也会这么想：当我们死后，世界依然存在。同样的太阳会继续照亮大地，同样的星球会继续转动，因为我们认为从开天辟地以来，它们一直都是如此。我们的孩子会继承这个地球。这都显示出我们对于不断流转的世间和一切现象是多么无知。
exit
===================================================
Current Tweet Length: 566
Exceed the Twitter length limit: true
===================================================
Now back to the cut tweet content
>>>>>>>「the 1th tweet content」
(1/3)这会是一条测试文案：以下语句摘自《正见》
潜意识中，我们期待自己会到达不再需要修理任何东西的境界。总有一天，我们会「从此过着快乐的生活」。我们深信「解决」的概念。好像我们所有经历的一切，到这一刻为止的生命，都只是在彩排。盛大的演出还没开始。对大多数的人来说，这种永无休止的处理
>>>>>>>「the 2th tweet content」
(2/3)和重新安排以及更新版本，就是生活的定义。事实上，我们是在等待生命开始。
我们往往也会这么想：当我们死后，世界依然存在。同样的太阳会继续照亮大地，同样的星球会继续转动，因为我们认为从开天辟地以来，它们一直都是如此。我们的孩子会继承这个地球。这都显示出我们对于不断流转的世间和一切
>>>>>>>「the 3th tweet content」
(3/3)现象是多么无知。
>>>>>>>「End of tweet cutting」
```
