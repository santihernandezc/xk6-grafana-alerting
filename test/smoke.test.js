// TODO: add some smoke tests

export const options = {
  thresholds: {
    checks: ["rate==1"],
  },
}

export default function () {
  greeting()
  base32()
  random()
}
