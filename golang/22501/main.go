package main

type Cell struct {
    next Cell
}

type Group struct {
    cells map[Cell]interface{}
}
