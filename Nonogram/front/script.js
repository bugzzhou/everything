const gridContainer = document.getElementById('gridContainer');
const headerRows = document.getElementById('headerRows');
const headerCols = document.getElementById('headerCols');
const clickButton = document.getElementById('clickButton');
const checkButton = document.getElementById('checkButton');
const clickCountElement = document.getElementById('clickCount');
let clickCount = 0;

let gridData = [];
let rows = [];
let cols = [];

// 配置常量
const IP = '192.168.3.138';
const PORT = '8686';
const BASE_URL = `http://${IP}:${PORT}/nono`;

// 提取 URL
const FILL_GRID_URL = `${BASE_URL}/fillGrid`;
const CHECK_GRID_URL = `${BASE_URL}/check`;

async function fetchGridData() {
  try {
    const response = await fetch(FILL_GRID_URL);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    gridData = data.grid;
    rows = data.rows;
    cols = data.cols;
    initGrid();
  } catch (error) {
    console.error('Error fetching grid data:', error);
  }
}

function initGrid() {
  headerCols.innerHTML = '';
  headerRows.innerHTML = '';
  gridContainer.innerHTML = '';

  // 初始化行头（原来的列头）
  for (let i = 0; i < rows.length; i++) {
    const rowHeader = document.createElement('div');
    rowHeader.classList.add('header-cell');
    rowHeader.style.display = 'grid';
    rowHeader.style.gridTemplateColumns = `repeat(${rows[i].length}, 1fr)`;
    for (let j = 0; j < rows[i].length; j++) {
      const cell = document.createElement('div');
      cell.textContent = rows[i][j];
      rowHeader.appendChild(cell);
    }
    headerRows.appendChild(rowHeader);
  }

  // 初始化列头（原来的行头）
  for (let i = 0; i < cols.length; i++) {
    const colHeader = document.createElement('div');
    colHeader.classList.add('header-cell');
    colHeader.style.display = 'grid';
    colHeader.style.gridTemplateRows = `repeat(${cols[i].length}, 1fr)`;
    for (let j = 0; j < cols[i].length; j++) {
      const cell = document.createElement('div');
      cell.textContent = cols[i][j];
      colHeader.appendChild(cell);
    }
    headerCols.appendChild(colHeader);
  }

  // 初始化网格
  for (let i = 0; i < gridData.length; i++) {
    for (let j = 0; j < gridData[i].length; j++) {
      const cell = document.createElement('div');
      cell.classList.add('cell');
      if (gridData[i][j] === 1) {
        cell.classList.add('black');
      }
      cell.addEventListener('click', () => toggleCell(i, j));
      gridContainer.appendChild(cell);
    }
  }
}

function toggleCell(row, col) {
  gridData[row][col] = 1 - gridData[row][col];
  const cell = gridContainer.children[row * gridData.length + col];
  cell.classList.toggle('black');
  updateClickCount();
}

function updateClickCount() {
  clickCount++;
  clickCountElement.textContent = clickCount;
}

async function checkGrid() {
  const payload = { gridData, rows, cols };
  try {
    const response = await fetch(CHECK_GRID_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    });
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const result = await response.json();
    if (result.status) {
      await fetchGridData();
    } else {
      alert('Check failed!');
    }
  } catch (error) {
    console.error('Error checking grid:', error);
  }
}

clickButton.addEventListener('click', () => updateClickCount());
checkButton.addEventListener('click', () => checkGrid());

fetchGridData();
