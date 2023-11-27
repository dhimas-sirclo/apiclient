import http from "k6/http";
import { sleep } from "k6";

export const options = {
  executor: "ramping-arrival-rate",
  stages: [
    {
      duration: "2h",
      target: 20000,
    },
  ],
};

export default function () {
  http.get("https://test.k6.io");
  sleep(1);
}
