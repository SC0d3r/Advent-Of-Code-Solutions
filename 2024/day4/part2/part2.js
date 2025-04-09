const fs = require('fs');
const path = require('path');

const findXMas = (i, j, twoD) => {
  const diag = findMAS(i, j, twoD, [[-1, -1], [1, 1]]);
  if (!diag) return '';
  const revDiag = findMAS(i, j, twoD, [[-1, 1], [1, -1]]);
  if (!revDiag) return '';
  return diag + revDiag;
};

const findMAS = (i, j, twoD, idx) => {
  if (idx.length !== 2) {
    throw new Error("this shouldn't happen");
  }

  let res = 'A';
  let key = `${i}${j}`;
  for (const id of idx) {
    key += `${i + id[0]}${j + id[1]}`;
    if (i + id[0] < 0 || i + id[0] >= twoD.length || j + id[1] < 0 || j + id[1] >= twoD[0].length) {
      continue;
    }
    if (res.length === 1) {
      res = twoD[i + id[0]][j + id[1]] + res;
    } else {
      res += twoD[i + id[0]][j + id[1]];
    }
  }
  if (res !== 'MAS' && res !== 'SAM') {
    return '';
  }
  return key;
};

const main = async () => {
  const startTime = Date.now();

  try {
    const data = fs.readFileSync(path.join(__dirname, '../inp1.txt'), 'utf-8');
    const lines = data.split('\n');
    const twoD = lines.map(line => line.trim().split(''));

    const results = new Set();
    const foundMap = new Map();
    let founds = 0;

    // Launch worker functions for each relevant cell concurrently
    const promises = [];

    for (let i = 0; i < twoD.length; i++) {
      for (let j = 0; j < twoD[i].length; j++) {
        if (twoD[i][j] === 'A') {
          promises.push(new Promise((resolve) => {
            const key = findXMas(i, j, twoD);
            if (key && !foundMap.has(key)) {
              foundMap.set(key, true);
              results.add(key);
              founds++;
            }
            resolve();
          }));
        }
      }
    }

    await Promise.all(promises);

    // console.log('foundMap', Array.from(foundMap.entries()));
    console.log('elapsed', Date.now() - startTime, "ms");
    console.log('founds', founds);
  } catch (err) {
    console.error('Error:', err);
  }
};

main();
