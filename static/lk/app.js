/*



*/
const options = {
    version: '9 2025-09-04',
    roleMap: {
        'admin':'Админ',
        'doctor':'Доктор',
        'client':'Клиент'
    }
}
var storage = {

}

const views = {
    header: (data={}) => (`
<nav id="mainMenu" class="mainMenu">
    <div class="mainMenuLogo">
        <img src="./img/openvetLogo.svg" alt="logo">
        <h2>OpenVET</h2>
    </div>
    ${data.clear ? '' : `

${currentUser.role == 'admin' ? `
    <h4></h4>
    <button onmousedown="pages.users({update:true})"><img src="./img/ico/people.svg" alt="ico">Пользователи</button>
    <button onmousedown="pages.report({update:true})"><img src="./img/ico/file-earmark-bar-graph-fill.svg" alt="ico">Отчёт</button>
    
    <h4></h4>
    <button onmousedown="pages.timesheets({update:true})"><img src="./img/ico/calendar2-week-fill.svg" alt="ico">Расписание</button>
    <button onmousedown="pages.records({update:true})"><img src="./img/ico/clipboard-heart-fill.svg" alt="ico">Записи</button>
    <button onmousedown="pages.clients({update:true})"><img src="./img/ico/person-vcard.svg" alt="ico">Клиенты</button>
    <button onmousedown="pages.animals({update:true})"><img src="./img/ico/bandaid-fill.svg" alt="ico">Животные</button>
    
    <h4></h4>
    <button onmousedown="pages.profile()"><img src="./img/ico/person.svg" alt="ico">Профиль</button>
    <button onmousedown="pages.help()"><img src="./img/ico/question-octagon.svg" alt="ico">Помощь</button>
    
    
`:''}
${currentUser.role == 'doctor' ? `
    <h4></h4>
    <button onmousedown="pages.timesheets({update:true})"><img src="./img/ico/calendar2-week-fill.svg" alt="ico">Расписание</button>
    <button onmousedown="pages.records({update:true})"><img src="./img/ico/clipboard-heart-fill.svg" alt="ico">Записи</button>
    <button onmousedown="pages.clients({update:true})"><img src="./img/ico/person-vcard.svg" alt="ico">Клиенты</button>
    <button onmousedown="pages.animals({update:true})"><img src="./img/ico/bandaid-fill.svg" alt="ico">Животные</button>
    <h4></h4>
    <button onmousedown="pages.profile()"><img src="./img/ico/person.svg" alt="ico">Профиль</button>
    <button onmousedown="pages.help()"><img src="./img/ico/question-octagon.svg" alt="ico">Помощь</button>
`:''}
${currentUser.role == 'client' ? `
    <h4></h4>
    <button onmousedown="location.href='tel:+79675552322'"><img src="./img/ico/telephone-fill.svg" alt="ico">Позвонить</button>
    <button onmousedown="pages.records({update:true})"><img src="./img/ico/clipboard-heart-fill.svg" alt="ico">Записи</button>
    <!--button onmousedown="pages.animals()"><img src="./img/ico/bandaid-fill.svg" alt="ico">Животные</button-->
    <h4></h4>
    <button onmousedown="pages.profile()"><img src="./img/ico/person.svg" alt="ico">Профиль</button>
    <button onmousedown="pages.help()"><img src="./img/ico/question-octagon.svg" alt="ico">Помощь</button>
`:''}



    
    
    `}
</nav>
    `),
    /*
     * расписание в виде календаря
     * Расписание для врачей привязанное к записям клиентов
     * 
     * Записи Описание
     * создано
     * отмена или подтверждение
     * выполнено
     * неявка
     * 
     * Клиенты
     * фио телефон ДоступКСистеме(привязку к пользователю) Описание
     * 
     * Животные
     * тип, кличка, дата рождения, привязка к клиенту, Описание
     * 
     * 
     * Пользователи
     * роль
     * фио
     * телефон
     * пароль
     * 
     * 
     * Отчёт
     * Клиентов всего
     * записей всего
     * записей всего отменённых
     * Успешных записей за текущий месяц
     * 
     * 
     * Профиль
     * фио, телефон, выйти
     * 
     * Помощь
     * Тут список частых вопросов и ответы на них + контакты
     * 
     * 
        <button onmousedown="pages.products()"><img src="./img/ico/shop.svg" alt="ico">Товары</button>
    <button onmousedown="pages.productsCategories()"><img src="./img/boxes.svg" alt="ico">Категории</button>
    <button onmousedown="pages.productsColors()"><img src="./img/palette.svg" alt="ico">Цвета товаров</button>
    <button onmousedown="pages.questions()"><img src="./img/question-octagon.svg" alt="ico">Вопросы</button>
    <button onmousedown="pages.texts()"><img src="./img/body-text.svg" alt="ico">Текста</button>
    */
    



    buttonRole : data => (`<button class="btn ${currentUser.role == data.code ? 'btn-mark' : ''}" onmousedown="currentUser.role='${data.code}';pages.help()">${data.name}</button>`),
}
let pages = {
    login: data => {
        utils.setText(`
<div class="loginBox">
    <h1 class="logo_text">OpenVET</h1>
    <form class="form1 block" id="formLogin" onsubmit="app.login(this);return false">
        <label class="name required">Номер телефона</label>
        <div class="value"><input 
        type="tel"
        required
        placeholder="+7 (___) ___-__-__"
        oninput="app.inputPhone(this)"
        name="phone" value="" maxlength="64"></div>
        <br>
        <label class="name required">Пароль</label>
        <div class="value"><input name="password" type="password" value="" maxlength="64"></div>
        <div class="buttons">
            <button class="btn btn-mark">Войти</button>
        </div>
        <div class="buttons">
            <button class="btn" type="button" onmousedown="notify.message('Восстановите через администратора')">Забыл пароль</button>
        </div>
    </form>
</div>
      `);
    },
    help: data => {
        utils.setText(` 
        ${views.header(data)}
        <div class="content" id="content">
            ${currentUser.role == '' ? '': `
            <div class="title block">
                <h1>Помощь</h1>
                <div class="text">
                    <p>
                    <b>О нас</b><br>
                    Лучшая вет клинника!
                    </p>
                    
                    <p><b>Контакты</b><br> 
                        Ярослав Ряднов (По всем вопросам)
                        <br><a href="https://t.me/openvet_contact" target="_blank" rel="nofollow">Telegram @openvet_contact</a>
						<br><a href="tel:+79675552322">Позвонить +7 (967) 555-23-22</a> (внутренний 100)
                    </p>
                    
                    <p>
                    <b>Ответы на частые вопросы</b><br>
                    А как кушать?
                    </p>
                </div>
            </div>
            `}
            <div class="title block" style="opacity: 0.5">
                <h1>Тест</h1>
                <div class="text">
                    <p>Тестовый режим для демонстрации возможностей каждой роли</p>
                    <p>Для продолжения выберите роль</p>
                </div>
                <div class="buttons">
                    ${Object.keys(options.roleMap).map(key => views.buttonRole({code:key, name:options.roleMap[key]})).join('')}
                </div>
            </div>
            
        </div>
        `)
    },

    users: data => {
        app.updateEntity(data, 'users')
        let items = storage.users || []

        if (data && data.filterRoleID) {
            items = items.filter(item => item.roleID == data.filterRoleID)
        }
        utils.setText(` 
${views.header(data)}
<div class="content" id="content">
    <div class="title block">
        <h1>Пользователи</h1>
        <div class="text">
            <p>Список пользователей, у которых есть доступ в личный кабинет</p>
            <p>Добавлять и изменять, менять роли - доступно только админу</p>
        </div>
        <div class="buttons">
            ${currentUser.role == 'admin' ? `<button class="btn btn-mark" onmousedown="pages.user()"><img src="./img/ico/plus-circle.svg" alt="ico">Добавить</button>`:''}
            <button class="btn" onmousedown="pages.users()">Все</button>
            <button class="btn${data && data.filterRoleID && data.filterRoleID == 4 ? ' btn-mark' : ''}" onmousedown="pages.users({filterRoleID:4})">Админы</button>
            <button class="btn${data && data.filterRoleID && data.filterRoleID == 3 ? ' btn-mark' : ''}" onmousedown="pages.users({filterRoleID:3})">Врачи</button>
            <button class="btn${data && data.filterRoleID && data.filterRoleID == 2 ? ' btn-mark' : ''}" onmousedown="pages.users({filterRoleID:2})">Клиенты</button>
        </div>
    </div>
    <div class="list block" style="grid-template-columns: 1fr 3fr 2fr 1fr 1fr;">
        <div class="line">
            <div class="grid-item grid-header">ID</div>
            <div class="grid-item grid-header">ФИО</div>
            <div class="grid-item grid-header">Телефон</div>
            <div class="grid-item grid-header">Роль</div>
            <div class="grid-item grid-header">Действия</div>
        </div>
        ${items.map(item => (`
        <div class="line">
            <div class="grid-item">${item.id}</div>
            <div class="grid-item">${item.fio}</div>
            <div class="grid-item">${item.phone}</div>
            <div class="grid-item">${storage.userRoles[item.roleID].name}</div>
            <div class="grid-item"><button class="btn" onmousedown="pages.user({id:${item.id}})">Править</button></div>
        </div>
        `)).join('')}
        
    </div>
</div>
        `)

    },
    user: data => {
        let user = { fio: '', roleID: 2, phone: '', password: '', description: '' };
        let isNew = true;

        if (data && data.id) {
            isNew = false
            const found = storage.users && storage.users.find(u => u.id === data.id)
            if (found) {
                user = found //{ ...found }
            } else {
                notify.err('Пользователь не найден')
                return;
            }
        }

        const roleOptions = Object.entries(storage.userRoles || {}).map(([id, role]) => ({
            id: Number(id),
            name: role.name,
            code: role.code
        }));

        utils.setText(`
${views.header(data)}
<div class="content" id="content">
  <div class="title block">
    <h1>${isNew ? 'Добавить пользователя' : 'Редактировать пользователя'}</h1>
    <div class="text">
      <p>Поля с * обязательны для заполнения.</p>
      <p>Только админ может управлять пользователями.</p>
    </div>
    ${isNew ? '' : `<div class="buttons">
        <button class="btn" onmousedown="app.deleteUser(${data.id})">Удалить</button>
    </div>`}
  </div>
  <form class="form1 block" onsubmit="app.saveUser(this, ${isNew ? 'true' : 'false'}); return false">
    <input type="hidden" name="id" value="${user.id || ''}">
    
    <label class="name required">ФИО</label>
    <div class="value">
      <input type="text" name="fio" value="${utils.escapeHtml(user.fio)}" required maxlength="255">
    </div>
    
    <label class="name required">Роль</label>
    <div class="value">
      <select name="roleID" required>
        ${roleOptions.map(r =>
            `<option value="${r.id}" ${user.roleID === r.id ? 'selected' : ''}>${r.name}</option>`
        ).join('')}
      </select>
    </div>
    
    <label class="name required">Телефон</label>
    <div class="value">
      <input type="tel" name="phone" value="${utils.escapeHtml(user.phone)}" 
        placeholder="+7 (___) ___-__-__"
        oninput="app.inputPhone(this)"
       maxlength="20">
    </div>
    
    <label class="name ${isNew ? 'required' : ''}">Пароль${isNew ? '' : ' (оставьте пустым, чтобы не менять)'}</label>
    <div class="value">
      <input type="password" name="password" ${isNew ? 'required' : ''} maxlength="64">
    </div>
    
    <label class="name">Описание</label>
    <div class="value">
      <textarea name="description" maxlength="500">${user.description ? utils.escapeHtml(user.description) : ''}</textarea>
    </div>
    
    <div class="buttons">
      <button class="btn btn-mark">${isNew ? 'Создать' : 'Сохранить'}</button>
      <button type="button" class="btn" onmousedown="pages.users()">Отмена</button>
    </div>
  </form>
</div>
`);
    },

    report: data => {
        app.updateEntity(data, 'report')

        let report = storage.report ? storage.report : {}
        utils.setText(`
    ${views.header(data)}
    <div class="content" id="content">
      <div class="title block">
        <h1>Отчёт</h1>
        <div class="text">
          <p>
            Клиентов всего: ${report.clientsTotal ? report.clientsTotal: '0'}<br>
            Записей всего: ${report.recordsTotal ? report.recordsTotal: '0'}<br>
            Записей всего отменённых: ${report.recordsCancelled ? report.recordsCancelled : '0'}<br>
          </p>
          <p>
            Доступен только админам
          </p>
        </div>
      </div>
    </div>
  `)
    },

    timesheets: data => {
        app.updateEntity(data, 'timesheets')
        app.updateEntity(data, 'users', () => {
            pages.timesheets()
        })
        app.updateEntity(data, 'records', () => {
            pages.timesheets()
        })
        if (!storage.users) storage.users = []

        let headDays = app.listMonthDays3()
        let times = app.listTimesDay().reverse()
        utils.setText(` 
        ${views.header(data)}
        <div class="content" id="content">
            <div class="title block">
                <h1>Расписание</h1>
                <div class="text">
                    <p>
                    1 приём = 30 минутам
                    </p><p>
                    Управление: месячным графиком только у админа, а управление дневным у админа и врача.
                    Каждый врач может управлять только своим дневным графиком.
                    </p>
                </div>
                <div class="buttons">
                    <button class="btn" onmousedown="pages.timesheets({update:true})">Обновить</button>
                </div>
            </div>
            <div class="title block">
                <div class="calendarMonth" id="calendarMonth">
                <div>title</div>${headDays.map(n => `<div>${n}</div>`).join('')}
                ${storage.users.filter(item => item.roleID == 3).map(user => `
                <button class="btn" onmousedown="pages.user({id:user.id})">${user.fio}</button>${headDays.map(n => `<button class="btn" onmousedown="">${n}</button>`).join('')}
                `).join('')}
</div>
                <div class="calendarDay" id="calendarDay">
                ${times.map(time => `<div>${time}</div>`).join('')}
</div>
            </div>
        </div>
        `)
    },

    records: data => {
        app.updateEntity(data, 'records', () => {

            storage.records.forEach(item => {
                item.dateTimeShow = `${item.dateTime.split('T')[0]}<br>${item.dateTime.split('T')[1].split('+')[0]}`
            })
            pages.records()
        })
        app.updateEntity(data, 'clients', () => {
            storage.clientsMapID = storage.clients.reduce((acc, client) => {
                acc[client.id] = client
                return acc
            }, {})
            pages.records()
        })
        app.updateEntity(data, 'animals', () => {

        })

        if (currentUser.role == "admin")
        app.updateEntity(data, 'users', () => {
            storage.usersMapID = storage.users.reduce((acc, user) => {
                acc[user.id] = user
                return acc
            }, {})
            pages.records()
        })
        let items = storage.records || []

        utils.setText(` 
${views.header(data)}
<div class="content" id="content">
    <div class="title block">
        <h1>Записи</h1>
        <div class="text">
            <p>Список записей на приём, которые включают в себя записи на приём и результаты</p>
            <p>Добавлять и просматривать могут все, менять только админы и врачи</p>
        </div>
        <div class="buttons">
            <button class="btn btn-mark" onmousedown="pages.record()"><img src="./img/ico/plus-circle.svg" alt="ico">Добавить</button>
        </div>
    </div>
    ${currentUser.role == 'admin'? `
    <div class="list block" style="grid-template-columns: 3vw 8vw 1fr 1fr 16vw 12vw;">
        <div class="line">
            <div class="grid-item grid-header">ID</div>
            <div class="grid-item grid-header">Статус</div>
            <div class="grid-item grid-header">Клиент</div>
            <div class="grid-item grid-header">Назначенный Врач</div>
            <div class="grid-item grid-header">Дата время визита</div>
            <div class="grid-item grid-header">Действия</div>
        </div>
         ${items.map(item => (`
        <div class="line">
            <div class="grid-item">${item.id}</div>
            <div class="grid-item">${storage.recordStatuses[item.statusID].name}</div>
            <div class="grid-item"><button class="btn" onmousedown="pages.client({id:${item.clientID}})">${storage.clientsMapID && storage.clientsMapID[item.clientID] ? storage.clientsMapID[item.clientID].fio : item.clientID}</button></div>
            <div class="grid-item"><button class="btn" onmousedown="pages.user({id:${item.userID}})">${storage.usersMapID && storage.usersMapID[item.userID] ? storage.usersMapID[item.userID].fio : item.userID}</button></div>
            <div class="grid-item">${item.dateTimeShow}</div>
            <div class="grid-item"><button class="btn" onmousedown="pages.record({id:${item.id}})">Править</button></div>
        </div>
        `)).join('')}
    </div>
    `:''}
        ${currentUser.role == 'doctor'? `
    <div class="list block" style="grid-template-columns: 1fr 1fr 16vw 2fr 1fr;">
        <div class="line">
            <div class="grid-item grid-header">ID</div>
            <div class="grid-item grid-header">Статус</div>
            <div class="grid-item grid-header">Клиент</div>
            <div class="grid-item grid-header">Дата время визита</div>
            <div class="grid-item grid-header">Действия</div>
        </div>
         ${items.map(item => (`
        <div class="line">
            <div class="grid-item">${item.id}</div>
            <div class="grid-item">${storage.recordStatuses[item.statusID].name}</div>
            <div class="grid-item"><button class="btn" onmousedown="pages.client({id:${item.clientID}})">${storage.clientsMapID && storage.clientsMapID[item.clientID] ? storage.clientsMapID[item.clientID].fio : item.clientID}</button></div>
            <div class="grid-item">${item.dateTimeShow}</div>
            <div class="grid-item"><button class="btn" onmousedown="pages.record({id:${item.id}})">Править</button></div>
        </div>
        `)).join('')}
    </div>
    `:''}
    ${currentUser.role == 'client' ? `
    <div class="list block" style="grid-template-columns: 1fr 1fr 2fr 2fr 2fr 1fr;">
        <div class="line">
            <div class="grid-item grid-header">ID</div>
            <div class="grid-item grid-header">Статус</div>
            <div class="grid-item grid-header">Дата время визита</div>
            <div class="grid-item grid-header">Подробнее</div>
        </div>
                ${items.map(item => (`
        <div class="line">
            <div class="grid-item">${item.id}</div>
            <div class="grid-item">${storage.recordStatuses[item.statusID].name}</div>
            <div class="grid-item">${item.dateTimeShow}</div>
            <div class="grid-item"><button class="btn" onmousedown="pages.record({id:${item.id}})">Подробнее</button></div>
        </div>
        `)).join('')}
    </div>
    `:''}
    
    
</div>
        `)

    },
    record: (data = {}) => {
        let record = {
            id: 0,
            clientID: 0,
            animalID: 0,
            userID: 0,
            statusID: 1, // по умолчанию "новый"
            dateTime: '',
            complaints: '',
            examination: '',
            recommendations: '',
        };
        let isNew = true;

        if (data && data.id && storage.records) {
            const found = storage.records.find(r => r.id === data.id);
            if (found) {
                record = found;
                isNew = false;
            } else {
                notify.err('Запись не найдена');
                return;
            }
        }

        // Подготавливаем данные
        const clients = storage.clients || [];
        const users = (storage.users || []).filter(u => {
            const role = storage.userRoles?.[u.roleID];
            return role?.code === 'doctor';
        });
        const statuses = Object.values(storage.recordStatuses || {}).map(s => ({
            id: s.id,
            name: s.name
        }));

        // Генерируем список животных (может обновляться динамически)
        const allAnimals = storage.animals || [];
        let animalsForClient = allAnimals.filter(a => a.clientID === record.clientID);

        utils.setText(`
${views.header(data)}
<div class="content" id="content">
  <div class="title block">
    <h1>${isNew ? 'Создать запись' : 'Редактировать запись'}</h1>
    <div class="text">
      <p>Поля с * обязательны для заполнения.</p>
      <p>Только доктор и админ могут управлять записями.</p>
    </div>
    ${!isNew ? `<div class="buttons">
      <button class="btn" onmousedown="app.deleteRecord(${record.id})">Удалить</button>
    </div>` : ''}
  </div>
  <form class="form1 block" onsubmit="app.saveRecord(this, ${isNew}); return false">
    <input type="hidden" name="id" value="${record.id}">
    
    <label class="name required">Клиент</label>
    <div class="value">
      <select name="clientID" required onchange="app.onClientChange(this)">
        <option value="">Выберите клиента</option>
        ${clients.map(c =>
            `<option value="${c.id}" ${record.clientID == c.id ? 'selected' : ''}>${utils.escapeHtml(c.fio)}</option>`
        ).join('')}
      </select>
    </div>

    <label class="name required">Животное</label>
    <div class="value">
      <select name="animalID" required id="animalSelect">
        <option value="">Выберите животное</option>
        ${animalsForClient.map(a =>
            `<option value="${a.id}" ${record.animalID == a.id ? 'selected' : ''}>${utils.escapeHtml(a.name)} (${storage.animalTypes?.[a.animalTypeID]?.name || '???'})</option>`
        ).join('')}
      </select>
    </div>

    <label class="name required">Врач</label>
    <div class="value">
      <select name="userID" required>
        <option value="">Выберите врача</option>
        ${users.map(u =>
            `<option value="${u.id}" ${record.userID == u.id ? 'selected' : ''}>${utils.escapeHtml(u.fio)}</option>`
        ).join('')}
      </select>
    </div>

    <label class="name required">Статус</label>
    <div class="value">
      <select name="statusID" required>
        ${statuses.map(s =>
            `<option value="${s.id}" ${record.statusID == s.id ? 'selected' : ''}>${utils.escapeHtml(s.name)}</option>`
        ).join('')}
      </select>
    </div>

    <label class="name required">Дата и время приёма</label>
    <div class="value">
      <input type="datetime-local" name="dateTime"
        value="${record.dateTime.slice(0, 16)}"
        required>
    </div>

    <label class="name">Жалобы</label>
    <div class="value">
      <textarea name="complaints" maxlength="500">${record.complaints ? utils.escapeHtml(record.complaints) : ''}</textarea>
    </div>

    <label class="name ">Осмотр</label>
    <div class="value">
      <textarea name="examination"  maxlength="1000">${record.examination ? utils.escapeHtml(record.examination) : ''}</textarea>
    </div>

    <label class="name ">Рекомендации</label>
    <div class="value">
      <textarea name="recommendations"  maxlength="1000">${record.recommendations ? utils.escapeHtml(record.recommendations) : ''}</textarea>
    </div>

    <div class="buttons">
      <button class="btn btn-mark">${isNew ? 'Создать запись' : 'Сохранить'}</button>
      <button type="button" class="btn" onmousedown="pages.records()">Отмена</button>
    </div>
  </form>
</div>
  `);
    },



    
    clients: data => {
        app.updateEntity(data, 'clients')
        let items = storage.clients || []
        utils.setText(` 
${views.header(data)}
<div class="content" id="content">
    <div class="title block">
        <h1>Клиенты</h1>
        <div class="text">
            <p>Список всех клиентов</p>
            <p>Добавлять, менять, просматривать только врачи и админы</p>
        </div>
        <div class="buttons">
            <button class="btn btn-mark" onmousedown="pages.client()"><img src="./img/ico/plus-circle.svg" alt="ico">Добавить клиента</button>
        </div>
    </div>
    <div class="list block" style="grid-template-columns: 1fr 3fr 2fr 1fr;">
        <div class="line">
            <div class="grid-item grid-header">ID</div>
            <div class="grid-item grid-header">ФИО Клиента</div>
            <div class="grid-item grid-header">Телефон</div>
            <div class="grid-item grid-header">Действия</div>
        </div>
        ${items.map(item => (`
        <div class="line">
            <div class="grid-item">${item.id}</div>
            <div class="grid-item">${item.fio}</div>
            <div class="grid-item">${item.phone}</div>
            <div class="grid-item"><button class="btn" onmousedown="pages.client({id:${item.id}})">Править</button></div>
        </div>
        `)).join('')}
        
    </div>
</div>
        `)
        
    },
    client: (data = {}) => {
        let client = {
            id: 0,
            family_name: '',
            name: '',
            middle_name: '',
            phone: '',
            email: '',
            birth_date: '',
            parent_id: 0,
            description: ''
        };
        let isNew = true;

        if (data && data.id && storage.clients) {
            const found = storage.clients.find(c => c.id === data.id);
            if (found) {
                client = found;
                isNew = false;
            } else {
                notify.err('Клиент не найден');
                return;
            }
        }

        utils.setText(`
${views.header(data)}
<div class="content" id="content">
  <div class="title block">
    <h1>${isNew ? 'Добавить клиента' : 'Редактировать клиента'}</h1>
    <div class="text">
      <p>Поля с * обязательны для заполнения.</p>
      <p>Только врачи и админы могут управлять клиентами.</p>
    </div>
    ${!isNew ? `<div class="buttons">
      <button class="btn" onmousedown="app.deleteClient(${client.id})">Удалить</button>
    </div>` : ''}
  </div>
  <form class="form1 block" onsubmit="app.saveClient(this, ${isNew}); return false">
    <input type="hidden" name="id" value="${client.id}">
    <label class="name required">ФИО</label>
    <div class="value">
      <input type="text" name="fio" value="${utils.escapeHtml(client.fio || '')}" required maxlength="255">
    </div>
    <label class="name required">Телефон</label>
    <div class="value">
      <input type="tel" name="phone" value="${utils.escapeHtml(client.phone || '')}"
        placeholder="+7 (___) ___-__-__"
        oninput="app.inputPhone(this)"
        required maxlength="20">
    </div>
     <label class="name">Дата рождения</label>
    <div class="value">
      <input type="text" name="birthDate" value="${client.birthDate ? new Date(client.birthDate).toISOString().split('T')[0] : ''}"
             placeholder="ГГГГ-ММ-ДД">
    </div>
    <label class="name">Описание</label>
    <div class="value">
      <textarea name="description" maxlength="500">${client.description ? utils.escapeHtml(client.description) : ''}</textarea>
    </div>
    <div class="buttons">
      <button class="btn btn-mark">${isNew ? 'Создать' : 'Сохранить'}</button>
      <button type="button" class="btn" onmousedown="pages.clients()">Отмена</button>
    </div>
  </form>
</div>
  `);
    },


    animals: data => {
        app.updateEntity(data, 'animals', () => {
            if (storage.animals) {}
            storage.animals.forEach(item => {
                item.birthDate = item.birthDate.split('T')[0]
            })
            pages.animals()
        })
        app.updateEntity(data, 'clients', () => {
            storage.clientsMapID = storage.clients.reduce((acc, client) => {
                acc[client.id] = client
                return acc
            }, {})
            pages.animals()
        })


        let items = storage.animals || []
        utils.setText(` 
${views.header(data)}
<div class="content" id="content">
    <div class="title block">
        <h1>Животные</h1>
        <div class="text">
            <p>Список всех животных</p>
            <p>Добавлять, менять, просматривать только врачи и админы</p>
        </div>
        <div class="buttons">
            <button class="btn btn-mark" onmousedown="pages.animal()"><img src="./img/ico/plus-circle.svg" alt="ico">Добавить живность</button>
        </div>
    </div>
    <div class="list block" style="grid-template-columns: 1fr 1fr 2fr 1fr 3fr 1fr;">
        <div class="line">
            <div class="grid-item grid-header">ID</div>
            <div class="grid-item grid-header">Тип</div>
            <div class="grid-item grid-header">Кличка</div>
            <div class="grid-item grid-header">Дата рождения</div>
            <div class="grid-item grid-header">Хозяин</div>
            <div class="grid-item grid-header">Действия</div>
        </div>
            ${items.map(item => (`
        <div class="line">
            <div class="grid-item">${item.id}</div>
            <div class="grid-item">${storage.animalTypes[item.animalTypeID].name}</div>
            <div class="grid-item">${item.name}</div>
            <div class="grid-item">${item.birthDate}</div>
            <div class="grid-item"><button class="btn" onmousedown="pages.client({id:${item.clientID}})">${storage.clientsMapID && storage.clientsMapID[item.clientID] ? storage.clientsMapID[item.clientID].fio : item.clientID}</button></div>
            <div class="grid-item"><button class="btn" onmousedown="pages.animal({id:${item.id}})">Править</button></div>
        </div>
        `)).join('')}
        
    </div>
</div>
        `)

    },
    animal: (data = {}) => {
        let animal = {
            id: 0,
            name: '',
            birthDate: '',
            chip: '',
            clientID: 0,
            animalTypeID: 0,
            description: '',
            parentID: 0,
        };
        let isNew = true;

        if (data && data.id && storage.animals) {
            const found = storage.animals.find(a => a.id === data.id);
            if (found) {
                animal = found;
                isNew = false;
            } else {
                notify.err('Животное не найдено');
                return;
            }
        }

        // Подготавливаем клиентов и типы
        const clients = storage.clients || [];
        const animalTypes = Object.values(storage.animalTypes || {}).map(t => ({
            id: t.id,
            name: t.name
        }));

        utils.setText(`
${views.header(data)}
<div class="content" id="content">
  <div class="title block">
    <h1>${isNew ? 'Добавить животное' : 'Редактировать животное'}</h1>
    <div class="text">
      <p>Поля с * обязательны для заполнения.</p>
      <p>Только врачи и админы могут управлять животными.</p>
    </div>
    ${!isNew ? `<div class="buttons">
      <button class="btn" onmousedown="app.deleteAnimal(${animal.id})">Удалить</button>
    </div>` : ''}
  </div>
  <form class="form1 block" onsubmit="app.saveAnimal(this, ${isNew}); return false">
    <input type="hidden" name="id" value="${animal.id}">
    <label class="name required">Кличка</label>
    <div class="value">
      <input type="text" name="name" value="${utils.escapeHtml(animal.name)}" required maxlength="255">
    </div>
    <label class="name">Дата рождения</label>
    <div class="value">
      <input type="text" name="birthDate" value="${animal.birthDate ? new Date(animal.birthDate).toISOString().split('T')[0] : ''}" placeholder="ГГГГ-ММ-ДД">
    </div>
    <label class="name">Чип (если есть)</label>
    <div class="value">
      <input type="text" name="chip" value="${animal.chip ? utils.escapeHtml(animal.chip) : ''}" maxlength="255">
    </div>
    <label class="name required">Тип</label>
    <div class="value">
      <select name="animalTypeID" required>
        <option value="">Выберите тип</option>
        ${animalTypes.map(t =>
            `<option value="${t.id}" ${animal.animalTypeID == t.id ? 'selected' : ''}>${utils.escapeHtml(t.name)}</option>`
        ).join('')}
      </select>
    </div>
    <label class="name required">Хозяин (клиент)</label>
    <div class="value">
      <select name="clientID" required>
        <option value="">Выберите клиента</option>
        ${clients.map(c =>
            `<option value="${c.id}" ${animal.clientID == c.id ? 'selected' : ''}>${utils.escapeHtml(c.fio)}</option>`
        ).join('')}
      </select>
    </div>
    <label class="name">Описание</label>
    <div class="value">
      <textarea name="description" maxlength="500">${animal.description ? utils.escapeHtml(animal.description) : ''}</textarea>
    </div>
    <div class="buttons">
      <button class="btn btn-mark">${isNew ? 'Создать' : 'Сохранить'}</button>
      <button type="button" class="btn" onmousedown="pages.animals()">Отмена</button>
    </div>
  </form>
</div>
  `);
    },




    profile:data => {
        if (!data) data = {};
        
        utils.setText(`
            ${views.header(data)}
            <div class="content" id="content">
                <div class="title block">
                    <h1>Профиль</h1>
                    <div class="text">
                        <p>
                        <b>ФИО:</b> ${currentUser.fio}<br>
                    <b>Роль:</b> ${options.roleMap[currentUser.role] ? options.roleMap[currentUser.role] : 'нет'}
                        
                        
</p>
                        <p>Менять только через админа.<br> Контакты в помощи</p>
                    </div>
                    <div class="buttons">
                        <button class="btn" onmousedown="app.logout()">Выйти</button>
                    </div>
                </div>
            </div>
        `);
    },
    
    

    
    
    
    
    // ---
    main: (data) => {
      if (!storage.user || !storage.user.id) {
        return pages.login();
      }
      if (!data || !data.stop)
        app.ajax("list", {}, (answer) => {
          if (answer.err) return app.showPopup(answer.err, true);
  
          storage.list = answer;
          pages.main({ stop: 1 });
        });

        // <p class="descr">При регистрации каждого клиента продавец получает 100 баллов.<br>(баллы выводятся через <a href="https://partners.OpenVET.ru/">partners.OpenVET.ru</a>)</p>
        utils.setText(` 
          ${app.addHeader()}
          <div class="main page">
              <h1 class="title">Зарегистрируйте клиента!</h1>
              <div class="tabs">
                  <button class="btn" onmousedown="pages.formAdd()">Добавить клиента</button>
                  <button class="btn btn-transp" onmousedown="pages.searchUserByBrCode()">
                    Оформить возврат
                  </button>
              </div>
              <div class="table">
                  <div class="table-header">
                      <p class="table-col">ID</p>
                      <p class="table-col">Имя клиента</p>
                      <p class="table-col">Телефон</p>
                      <p class="table-col">Код продукта</p>
                      <p class="table-col">Устройство</p>
                  </div>
                  ${storage.list ? storage.list.map(item => `
                      <div class="user-list-item">
                          <div class="table-row">
                            <p class="table-label">ID:</p>
                            <p class="table-value">${item.id}</p>
                          </div>
                          <div class="table-row">
                            <p class="table-label">Имя клиента:</p>
                            <p class="table-value">${item.family_name} ${item.name}</p>
                          </div>
                          <div class="table-row">
                            <p class="table-label">Телефон:</p>
                            <p class="table-value">${item.phone}</p>
                          </div>
                          <div class="table-row">
                            <p class="table-label">Код продукта:</p>
                            <p class="table-value">${item.product_code}</p>
                          </div>
                          <div class="table-row">
                            <p class="table-label">Устройство:</p>
                            <p class="table-value ${item.product_name && item.product_name.split(' ').length > 1 ? item.product_name.split(' ')[1].toLowerCase() : ''}">${item.product_name && item.product_name.split(' ').length > 1 ? item.product_name.split(' ')[1] : ''}</p>
                          </div>
                      </div>`).join("") : ""}
              </div>
              <div id="popup">
                  <span id="close-btn" onmousedown="app.closePopup()"><svg id="closeSvg" width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M7.25367 18C6.93501 18 6.61635 17.8826 6.36478 17.6312C5.87841 17.145 5.87841 16.3402 6.36478 15.854L15.8574 6.36465C16.3438 5.87845 17.1488 5.87845 17.6352 6.36465C18.1216 6.85086 18.1216 7.65561 17.6352 8.14181L8.14256 17.6312C7.90776 17.8826 7.57233 18 7.25367 18Z" fill="#A2E9C1"/>
                      <path d="M16.7463 18C16.4277 18 16.109 17.8826 15.8574 17.6312L6.36478 8.14181C5.87841 7.65561 5.87841 6.85086 6.36478 6.36465C6.85115 5.87845 7.65618 5.87845 8.14256 6.36465L17.6352 15.854C18.1216 16.3402 18.1216 17.145 17.6352 17.6312C17.3836 17.8826 17.065 18 16.7463 18Z" fill="#A2E9C1"/>
                  </svg></span>
                  <p id="popup-message"></p>
              </div>
          </div>
      `);
      window.scrollTo(0,0);
    },
    

}

const app = {


    main: () => {
        notify.start()
        app.initData()
    },
    initData: () => {
        ajax.getJson('/api/init', answer => {
            app.inputInnitData(answer)
            if (!storage.user)
                pages.login()
        })
    },
    inputInnitData: (answer) => {
        console.log(['init', answer])
        if (!answer || !answer.data || !answer.data.user) return;

        Object.assign(storage, answer.data)

        let user = answer.data.user
        currentUser.id = user.id
        currentUser.fio = user.fio
        currentUser.role = answer.data.userRoles && answer.data.userRoles[user.roleID] ? answer.data.userRoles[user.roleID].code : ''
        pages.help()

    },
    login: (form) => {
        const data = utils.formData(form)
        if (data.phone) data.phone = data.phone.replace(/\D/g, '')

        ajax.json('/api/login', data, answer => {
            app.inputInnitData(answer)

            if (!answer || !answer.ok)
                notify.err('Неверный логин или пароль')
        })
    },
    logout: () => {
        storage = {}
        ajax.getJson('/api/logout', answer => {
            pages.login()
        })
    },
    inputPhone: input => {
        input.value = utils.formatPhone(input.value)
    },
    // Обновляет ссписок перед загрузкой их на страницу
    updateEntity: (data, name='users', callback) => {
        if (!data || !data.update) return;

        ajax.getJson(`/api/${name}`, answer => {
            if (!answer) return;
            if (answer.error) {
                notify.err(answer.error)
                return;
            }
            storage[name] = answer[name] ? answer[name] : []
            if (storage[name] && Array.isArray(storage[name])) storage[name].reverse()

            if (typeof callback == 'function') {
                callback()
                return;
            }
            pages[name]()
        })
    },

    saveUser: (form, isNew) => {
        const formData = new FormData(form)
        const data = {
            id: formData.get('id') || undefined,
            fio: formData.get('fio').trim(),
            roleID: parseInt(formData.get('roleID'), 10),
            phone: formData.get('phone').trim().replace(/\D/g, ''),
            description: formData.get('description').trim() || null
        };

        // Пароль отправляем ТОЛЬКО если он указан
        const password = formData.get('password').trim()
        if (password) {
            data.password = password
        } else if (isNew) {
            notify.err('Пароль обязателен при создании')
            return;
        }


        if (!isNew && data.id) {
            data.id = parseInt(data.id, 10)
            data.ajaxMethod = 'PUT'
        }

        const url = isNew ? '/api/users' : `/api/users/${data.id}`

        ajax.json(url, data, answer => {
            if (answer && answer.ok) {
                notify.message(isNew ? 'Пользователь создан' : 'Пользователь обновлён')
                pages.users({ update: true })
            } else {
                notify.err(answer?.error || 'Ошибка сохранения')
            }
        })
    },
    deleteUser: (id) => {
       if (prompt('Для подтверждения введите "удалить"', '') != 'удалить') return;

        ajax.json(`/api/users/${id}`, {ajaxMethod: 'DELETE'}, answer => {
            if (answer && answer.ok) {
                notify.message('Пользователь удалён')
                pages.users({ update: true })
            } else {
                notify.err(answer?.error || 'Ошибка удаления')
            }
        })
    },

    saveClient: (form, isNew) => {
        const formData = new FormData(form);
        const data = {
            fio: formData.get('fio').trim(),
            phone: formData.get('phone').trim().replace(/\D/g, ''),
            description: formData.get('description').trim() || null,
            birthDate: utils.dateRFC3339(formData.get('birthDate').trim()),
        };

        if (!isNew) {
            //data.id = parseInt(formData.get('id'), 10)
            data.ajaxMethod = 'PUT'
        }
        const url = isNew ? '/api/clients' : `/api/clients/${formData.get('id')}`;

        ajax.json(url, data, function(answer) {
            if (answer && answer.ok) {
                notify.message(isNew ? 'Клиент создан' : 'Клиент обновлён');
                pages.clients({ update: true });
            } else {
                notify.err(answer?.error || 'Ошибка сохранения клиента');
            }
        });
    },

    deleteClient: id => {
        if (prompt('Для подтверждения введите "удалить"', '') !== 'удалить') return;
        ajax.json(`/api/clients/${id}`, { ajaxMethod: 'DELETE' }, function(answer) {
            if (answer && answer.ok) {
                notify.message('Клиент удалён');
                pages.clients({ update: true });
            } else {
                notify.err(answer?.error || 'Ошибка удаления клиента');
            }
        });
    },

    saveAnimal: function(form, isNew) {
        const formData = new FormData(form);
        const data = {
            name: formData.get('name').trim(),
            clientID: parseInt(formData.get('clientID'), 10),
            animalTypeID: parseInt(formData.get('animalTypeID'), 10),
        };

        const chip = formData.get('chip').trim();
        if (chip) data.chip = chip;

        const description = formData.get('description').trim();
        if (description) data.description = description;

        const birthDate = formData.get('birthDate').trim();
        if (birthDate) data.birthDate = utils.dateRFC3339(birthDate);

        if (!data.name || !data.clientID || !data.animalTypeID) {
            notify.err('Заполните все обязательные поля');
            return;
        }

        if (!isNew) {
            //data.id = parseInt(formData.get('id'), 10)
            data.ajaxMethod = 'PUT'
        }
        const url = isNew ? '/api/animals' : `/api/animals/${formData.get('id')}`;

        ajax.json(url, data, function(answer) {
            if (answer && answer.ok) {
                notify.message(isNew ? 'Животное создано' : 'Животное обновлено');
                pages.animals({ update: true });
            } else {
                notify.err(answer?.error || 'Ошибка сохранения животного');
            }
        });
    },

    deleteAnimal: function(id) {
        if (prompt('Для подтверждения введите "удалить"', '') !== 'удалить') return;
        ajax.json(`/api/animals/${id}`, { ajaxMethod: 'DELETE' }, function(answer) {
            if (answer && answer.ok) {
                notify.message('Животное удалено');
                pages.animals({ update: true });
            } else {
                notify.err(answer?.error || 'Ошибка удаления животного');
            }
        });
    },

    // Обновление списка животных при смене клиента
    onClientChange: function(select) {
        const clientId = parseInt(select.value, 10);
        const animalSelect = document.getElementById('animalSelect');
        if (!animalSelect) return;

        const allAnimals = storage.animals || [];
        const filtered = clientId
            ? allAnimals.filter(a => a.clientID === clientId)
            : [];

        animalSelect.innerHTML = '<option value="">Выберите животное</option>' +
            filtered.map(a =>
                `<option value="${a.id}">${utils.escapeHtml(a.name)} (${storage.animalTypes?.[a.animalTypeID]?.name || '???'})</option>`
            ).join('');
    },

// Сохранение записи
    saveRecord: function(form, isNew) {
        const data = {
            clientID: parseInt(form.clientID.value, 10),
            animalID: parseInt(form.animalID.value, 10),
            userID: parseInt(form.userID.value, 10),
            statusID: parseInt(form.statusID.value, 10),
            dateTime: utils.dateRFC3339(form.dateTime.value),
            examination: form.examination.value.trim(),
            recommendations: form.recommendations.value.trim(),
        };

        const complaints = form.complaints.value.trim();
        if (complaints) data.complaints = complaints;

        // Валидация
        if (!data.clientID || !data.animalID || !data.userID ||
            !data.statusID || !data.dateTime) {
            notify.err('Заполните все обязательные поля');
            return;
        }


        if (!isNew) {
            //data.id = parseInt(formData.get('id'), 10)
            data.ajaxMethod = 'PUT'
        }
        const url = isNew ? '/api/records' : `/api/records/${form.id.value}`;


        ajax.json(url, data, function(answer) {
            if (answer && answer.ok) {
                notify.message(isNew ? 'Запись создана' : 'Запись обновлена');
                pages.records({ update: true });
            } else {
                notify.err(answer?.error || 'Ошибка сохранения записи');
            }
        });
    },

// Удаление записи
    deleteRecord: function(id) {
        if (prompt('Для подтверждения введите "удалить"', '') !== 'удалить') return;
        ajax.json(`/api/records/${id}`, { ajaxMethod: 'DELETE' }, function(answer) {
            if (answer && answer.ok) {
                notify.message('Запись удалена');
                pages.records({ update: true });
            } else {
                notify.err(answer?.error || 'Ошибка удаления записи');
            }
        });
    },


    listMonthDays3:() => {
        const backDays = 3
        const nextDays = 14

        const totalDays = backDays + nextDays + 1;
        const result = new Array(totalDays);
        const now = new Date()

        now.setDate(now.getDate() - backDays);

        for (let i = 0; i < totalDays; i++) {
            result[i] = now.getDate()
            now.setDate(now.getDate() + 1)
        }

        return result
    },
    listTimesDay:() => {
        const times = [];
        for (let hour = 8; hour <= 20; hour++) {
            times.push(`${String(hour).padStart(2, '0')}:00`)
            times.push(`${String(hour).padStart(2, '0')}:30`)
        }

        return times
    }
    
}

const currentUser = {
    id: 0,
    fio: 'Test',
    role: '', // client || doctor || admin
}

// error to log
{
    /*
window.addEventListener('error', function(event) {
    let data = {
        'userID': 0,
        'message': event.message,
        'file': event.filename,
        'line': event.lineno + ':' + event.colno,
        'userAgent': navigator.userAgent,
        'url':location.href,
        'timeAfterStartPage': 0
    }
    let formData = new FormData()
    for (let key in data)
        formData.append(key, data[key])
    fetch('/api/addError', {
        method: 'POST',
        body: formData
    })
})
     */
}


const tests = {
    all: () => {
        ajax.json('/api/login', {phone:'10001110011', password:'ps'})
        ajax.json('/api/current_user', {})
        ajax.getJson('/api/init')


        ajax.getJson('/api/users')


    },
}
app.main()
