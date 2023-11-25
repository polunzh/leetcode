const sleep = require('./index');

describe('sleep', () => {
  test('sleeps for 1 second', async () => {
    const start = Date.now();
    await sleep(100);
    const end = Date.now();
    expect(end - start).toBeGreaterThanOrEqual(100);
  });
});
