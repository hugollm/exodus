# Exodus

Library that helps the writing of distributed genetic algorithms.


## Install

As any other Go library, you can get the code with:

    go get github.com/hugollm/exodus


## Usage Example

The following example takes a population of 10 sized vectors of zeros and ones and evolves it until it finds an individual that's all ones.

```go
package main

import "fmt"
import "github.com/hugollm/exodus"

func main() {
    search := Search{
        IndividualSize: 10,
        PopulationSize: 30,
        CrossoverRate: 0.6,
        MutationRate: 0.01,
        NewGene: newGene,
        Fitness: fitness,
        OnGeneration: onGeneration,
    }
    search.Start()
}

func newGene() float64 {
    return float64(RandomInt(0, 1))
}

func fitness(genome []float64) float64 {
    fit := 0.0
    for i := 0; i < len(genome); i++ {
        fit += genome[i]
    }
    return float64(fit)
}

func onGeneration(search *Search) {
    if (search.Best.Fitness == 10) {
        fmt.Println(search.Best.Genome)
        search.Stop()
    }
}
```


## Distributing Over Machines

By default the library will disbribute the fitness load over the machine processors in a master-slave distributed model.

However, the library is also capable of distributing the algorithm over other machines, in an island distributed model. To leverage that, you'll need to:

* Add a `MigrationRate` attribute to your search configuration.
* Spawn a server (main island), just running the code normally.
* Spawn clients (islands), running the code with an `EXODUS_SERVER` environment variable set to the url of the server (port 2345).

Example - search parameter:

```go
search := Search{
    // ...
    MigrationRate: 0.005,
    // ...
}
```

Example - spawning a server:

    mysearch

Example - spawning a client:

    EXODUS_SERVER='http://localhost:2345' mysearch
