# Scores

This is a quiz score tracker designed to keep track of the scores in a
quiz where each team in the quiz contributes a round and therefore
cannot receive a score in that round.

When doing this with my family twice a week during quarantine I
eventually decided that the method we were using to assign a score to
the team which made the round, which was giving them the highest score
of the other teams, was unfair and we should look to other methods such
as using the median or mean score of the other teams.

This repo provides a client/server system which allows for tracking of
scores according to each of the metrics mentioned above.

## Building
To build the server
```
bazel build //cmd/scores:scores
```

To build the client
```
bazel build //cmd/client:client
```

## Running

To run the server normally
```
bazel run //cmd/scores:scores
```

To run the server in a docker container
```
bazel run //cmd/scores:scores_container
```

To run the client
```
bazel run //cmd/client:client
```
