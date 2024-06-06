// 新增数据 函数
function addRow() {
    
    var table = document.getElementById('table');
    console.log(table);
    var length = table.rows.length;
    // 插入行节点
    var newRow = table.insertRow(length)
    console.log(newRow);
    var nameCol = newRow.insertCell(0)
    var phoneCol = newRow.insertCell(1)
    var actionCol = newRow.insertCell(2)
    console.log(newRow);
    // 修改节点文本内容
    nameCol.innerHTML = '未命名';
    phoneCol.innerHTML = '无联系方式';
    actionCol.innerHTML = '<button>编辑</button><button>删除</button>';

} 

