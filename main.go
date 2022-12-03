package main

import (
    "github.com/zhengow/vngo/model"
    "time"

    "github.com/zhengow/vngo/consts"
    "github.com/zhengow/vngo/database"
    "github.com/zhengow/vngo/engine"
    "github.com/zhengow/vngo/strategy"
)

func getSymbols(symbols []string, exchange consts.Exchange) []*model.Symbol {
    res := make([]*model.Symbol, 0)
    for _, symbol := range symbols {
        res = append(res, model.NewSymbol(symbol, exchange))
    }
    return res
}

func main() {
    b := engine.NewBacktestingEngine()
    symbols := getSymbols([]string{"BTCDOMUSDT"}, consts.ExchangeEnum.BINANCE)
    startDate := time.Date(2022, 7, 1, 0, 0, 0, 0, time.Local)
    endDate := time.Date(2022, 8, 1, 0, 0, 0, 0, time.Local)
    b.SetParameters(symbols, consts.IntervalEnum.MINUTE, startDate, endDate, nil, 10000)
    b.AddStrategy(&strategy.MyStrategy{}, nil)
    database.NewMysql()
    b.LoadData()
    b.RunBacktesting()
    b.CalculateResult()
    b.ShowChart()
}
