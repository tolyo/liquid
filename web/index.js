import {
  baseRates,
  sendFxRateUpdate,
  getRandomInterval,
  currencies,
} from "./mockrate.js";

angular
  .module("test", [])
  .controller(
    "BalanceController",
    class BalanceController {
      constructor($http, $rootScope) {
        this.$http = $http;
        this.balances = [];
        $rootScope.$on("updateBalance", () => {
          this.updateBalance();
        });
      }

      $onInit() {
        this.updateBalance();
      }

      updateBalance() {
        this.$http.get("/master").then((res) => {
          this.balances = res.data.data;
        });
      }
    },
  )

  .controller(
    "RateController",
    class RateController {
      constructor($http, $interval) {
        this.$http = $http;
        this.rates = baseRates;
        $interval(() => {
          sendFxRateUpdate()
        }, getRandomInterval())
      }
    },
  )

  .controller(
    "PaymentController",
    class PaymentController {
      constructor($http, $rootScope, $timeout) {
        this.$http = $http;
        this.$rootScope = $rootScope;
        this.$timeout = $timeout;
        this.currencies = currencies;
        this.rates = baseRates;
        this.balances = [];
        this.user = "USER_1";
        this.amountSell = 0;
        this.amountBuy = 0;
        this.sellCurrency = "USD";
        this.buyCurrency = "EUR";
        this.recalculateSellAmount();
      }

      $onInit() {
        this.updateBalance();
      }

      updateBalance() {
        this.$http.get("/balances/" + this.user).then((res) => {
          this.balances = res.data.data;
        });
      }

      selectSellCurrency() {
        if (this.buyCurrency) {
          this.buyCurrency = undefined;
        }
      }

      availableCurrencies() {
        return this.currencies.filter((x) => x !== this.sellCurrency);
      }

      recalculateBuyAmount() {
        this.rate = this.rates[`${this.buyCurrency}/${this.sellCurrency}`];
        const amountBuy = (this.amountSell * 100) / this.rate;
        if (isNaN(amountBuy)) {
          this.amountBuy = 0;
        } else {
          this.amountBuy = Math.round(amountBuy) / 100;
        }
      }

      recalculateSellAmount() {
        this.rate = this.rates[`${this.buyCurrency}/${this.sellCurrency}`];
        const amountSell = (this.amountBuy * 100) * this.rate;
        if (isNaN(amountSell)) {
          this.amountSell = 0;
        } else {
          this.amountSell = Math.round(amountSell) / 100;
        }
      }

      submit() {
        this.$http
          .post("/transfer", {
            amount: this.amountSell * 100,
            currencySell: this.sellCurrency,
            currencyBuy: this.buyCurrency,
            externalId: this.user,
          })
          .then((res) => {
            this.updateBalance();
            this.$rootScope.$broadcast("updateBalance", {});
          })
          .catch((e) => {
            this.error = true;
            this.errorDescription = e.data.data;
            this.$timeout(() => {
              this.error = false
            }, 3000)
          }); 
      }
    },
  )

  .filter("currency", () => {
    return (input) => {
      const res = parseInt(input) / 100;
      return res.toFixed(2);
    };
  });
