describe('sleep', () => {
  test('sleeps for 1 second', async () => {
    const sleep = require('./index');
    const start = Date.now();
    await sleep(100);
    const end = Date.now();
    expect(end - start).toBeGreaterThanOrEqual(100);
  });
});
