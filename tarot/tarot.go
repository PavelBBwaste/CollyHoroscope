package tarot

import (
	"math/rand"
)

type TarotCard struct {
	Name        string
	Description string
	ImagePath   string
}

var tarotCards = []TarotCard{
	// Старшие Арканы
	{"Шут", "Начало, спонтанность, свобода, путешествие, открытие нового пути. Значение: Шут указывает на начало нового путешествия или этапа в жизни. Он символизирует смелость следовать за мечтами, принимая неожиданности и не боясь рискнуть.", "images/tarot/1.png"},
	{"Маг", "Воля, мастерство, ресурсы, инициатива, создание. Значение: Маг олицетворяет силу намерения и способности человека достичь целей. Он показывает, что время использовать свои навыки и действовать.", "images/tarot/2.png"},
	{"Жрица", "Интуиция, тайна, подсознание, внутреннее знание, духовное развитие. Значение: Жрица связана с глубоким пониманием вещей и призывает полагаться на интуицию. Это время для саморефлексии и поиска ответов внутри себя.", "images/tarot/3.png"},
	{"Императрица", "Плодородие, изобилие, материнство, творчество, природа. Значение: Императрица говорит о щедрости, процветании и заботе. Она указывает на творческую энергию и благоприятное время для создания чего-то нового.", "images/tarot/4.png"},
	{"Император", "Власть, контроль, структура, порядок, ответственность. Значение: Император символизирует порядок и силу. Он указывает на необходимость проявить лидерство и взять на себя ответственность за происходящее.", "images/tarot/5.png"},
	{"Жрец (Папа)", "Традиции, духовное руководство, религия, мудрость, наставничество. Значение: Жрец подчеркивает важность следования традициям и уважения к духовным учениям. Это карта поиска мудрости у авторитетных личностей.", "images/tarot/6.png"},
	{"Влюбленные", "Любовь, выбор, гармония, партнерство, привязанность. Значение: Влюбленные символизируют важный выбор или возможность, связанную с сердечными делами и взаимоотношениями.", "images/tarot/7.png"},
	{"Колесница", "Триумф, движение вперед, победа, сила воли, путешествие. Значение: Колесница указывает на стремительное движение к цели. Это карта амбиций, контроля и победы над обстоятельствами.", "images/tarot/8.png"},
	{"Сила", "Внутренняя сила, мужество, терпение, самообладание, сострадание. Значение: Карта Силы указывает на мощный внутренний потенциал и смелость. Это время проявить терпение и мудрость.", "images/tarot/9.png"},
	{"Отшельник", "Одиночество, самопознание, мудрость, внутренний поиск, отстраненность. Значение: Отшельник символизирует необходимость уединения для поиска истины. Это карта саморефлексии и внутреннего поиска.", "images/tarot/10.png"},
	{"Колесо Фортуны", "Судьба, изменение, удача, циклы, неожиданности. Значение: Колесо Фортуны символизирует перемену обстоятельств и циклическую природу жизни. Это карта судьбоносных событий и шансов.", "images/tarot/11.png"},
	{"Справедливость", "Справедливость, правда, ответственность, закон, честность. Значение: Карта Справедливости означает объективность и честное отношение к ситуации. Важно принимать решения взвешенно.", "images/tarot/12.png"},
	{"Повешенный", "Жертвенность, пауза, изменение перспективы, принятие, отказ. Значение: Повешенный указывает на период ожидания и осознания. Это время для переосмысления и выхода за рамки обыденного.", "images/tarot/13.png"},
	{"Смерть", "Трансформация, окончание, перерождение, изменения, конец цикла. Значение: Смерть указывает на завершение одного этапа и начало нового. Это карта перемен и обновления.", "images/tarot/14.png"},
	{"Умеренность", "Баланс, гармония, умеренность, терпение, адаптация. Значение: Умеренность символизирует гармонию и важность соблюдения равновесия. Это время для примирения и нахождения баланса.", "images/tarot/15.png"},
	{"Дьявол", "Искушение, привязанность, зависимость, ограничение, иллюзия. Значение: Дьявол символизирует зависимости и привязанности, которые мешают свободе. Это предупреждение о самоограничениях и ловушках.", "images/tarot/16.png"},
	{"Башня", "Разрушение, перемены, кризис, катастрофа, шок. Значение: Башня указывает на внезапные и резкие изменения. Она предупреждает о необходимости отпустить старое и быть готовым к переменам.", "images/tarot/17.png"},
	{"Звезда", "Надежда, вдохновение, вера, исцеление, духовное руководство. Значение: Звезда символизирует свет надежды и вдохновения. Она указывает на восстановление после трудностей и оптимизм.", "images/tarot/18.png"},
	{"Луна", "Иллюзия, страхи, интуиция, подсознание, загадочность. Значение: Луна отражает скрытые чувства, интуицию и возможные заблуждения. Важно доверять интуиции и не бояться неизвестности.", "images/tarot/19.png"},
	{"Солнце", "Радость, успех, процветание, жизненная сила, оптимизм. Значение: Солнце — это карта успеха и радости. Она указывает на время счастья, роста и процветания.", "images/tarot/20.png"},
	{"Суд", "Пробуждение, самооценка, возрождение, судьба, кармическое воздаяние. Значение: Суд говорит о перерождении и осознании. Это период самооценки и понимания последствий своих поступков.", "images/tarot/21.png"},
	{"Мир", "Завершение, достижение, гармония, завершение цикла, удовлетворение. Значение: Мир символизирует исполнение и завершение цикла. Это карта гармонии, успеха и удовлетворения.", "images/tarot/22.png"},

	// Масть Жезлов
	{"Туз Жезлов", "Вдохновение, новые начинания, амбиции. Значение: Туз Жезлов — карта, символизирующая начало творческого или делового пути. Это время для новых идей и энергии.", "images/tarot/23.png"},
	{"Двойка Жезлов", "Планирование, принятие решений, перспектива. Значение: Двойка Жезлов говорит о периоде анализа и выбора. Это карта стратегического планирования и оценивания доступных возможностей.", "images/tarot/24.png"},
	{"Тройка Жезлов", "Расширение, дальновидность, путешествие. Значение: Тройка Жезлов символизирует время для роста и движения вперёд. Это карта уверенности в своём пути и готовности к новым горизонтам.", "images/tarot/25.png"},
	{"Четверка Жезлов", "Празднование, стабильность, достижение. Значение: Четверка Жезлов говорит о моменте радости и отдыха после упорного труда. Это карта стабильности и радостных событий.", "images/tarot/26.png"},
	{"Пятерка Жезлов", "Конкуренция, борьба, конфликт. Значение: Пятерка Жезлов указывает на напряжение и борьбу. Это карта соперничества и испытаний, требующих упорства.", "images/tarot/27.png"},
	{"Шестерка Жезлов", "Победа, признание, триумф. Значение: Шестерка Жезлов символизирует успех и признание достижений. Это карта уверенности в своих силах и ощущения победы.", "images/tarot/28.png"},
	{"Семерка Жезлов", "Защита, уверенность, преодоление препятствий. Значение: Семерка Жезлов говорит о необходимости защищать свои убеждения и преодолевать вызовы. Это карта смелости и настойчивости.", "images/tarot/29.png"},
	{"Восьмерка Жезлов", "Движение, ускорение, прогресс. Значение: Восьмерка Жезлов символизирует стремительное движение и прогресс. Это карта быстрых перемен и возможности для действий.", "images/tarot/30.png"},
	{"Девятка Жезлов", "Настойчивость, защита, готовность к борьбе. Значение: Девятка Жезлов указывает на необходимость сохранять бдительность. Это карта силы и готовности преодолеть трудности.", "images/tarot/31.png"},
	{"Десятка Жезлов", "Ответственность, бремя, давление. Значение: Десятка Жезлов символизирует тяжелое бремя и чувство давления. Это карта трудностей, но и силы, чтобы их преодолеть.", "images/tarot/32.png"},
	{"Паж Жезлов", "Исследование, энтузиазм, творчество. Значение: Паж Жезлов указывает на любопытство и открытость новому. Это карта начинаний и энтузиазма.", "images/tarot/33.png"},
	{"Рыцарь Жезлов", "Стремление, авантюризм, динамика. Значение: Рыцарь Жезлов олицетворяет энергию и стремление к приключениям. Это карта движения к цели с полной отдачей.", "images/tarot/34.png"},
	{"Королева Жезлов", "Страсть, уверенность, харизма. Значение: Королева Жезлов указывает на силу и харизму. Это карта вдохновения и теплой энергии.", "images/tarot/35.png"},
	{"Король Жезлов", "Лидерство, предпринимательство, вдохновение. Значение: Король Жезлов символизирует уверенное лидерство и готовность действовать. Это карта силы и уверенности.", "images/tarot/36.png"},

	// Масть Кубков
	{"Туз Кубков", "Любовь, начало новых отношений, духовное удовлетворение. Значение: Туз Кубков указывает на начало эмоционального пути, наполненного любовью и гармонией. Это карта новых отношений и радости.", "images/tarot/37.png"},
	{"Двойка Кубков", "Партнерство, привязанность, союз. Значение: Двойка Кубков символизирует гармоничные отношения и крепкие связи. Это карта любви, привязанности и союза.", "images/tarot/38.png"},
	{"Тройка Кубков", "Дружба, радость, празднование. Значение: Тройка Кубков олицетворяет радость, дружбу и время для празднования. Это карта позитивных эмоций и радостных событий.", "images/tarot/39.png"},
	{"Четверка Кубков", "Апатия, разочарование, самоанализ. Значение: Четверка Кубков говорит о состоянии разочарования или апатии. Это карта необходимости пересмотра приоритетов.", "images/tarot/40.png"},
	{"Пятерка Кубков", "Потеря, сожаление, печаль. Значение: Пятерка Кубков символизирует печаль и утрату, но также дает надежду на обновление.", "images/tarot/41.png"},
	{"Шестерка Кубков", "Воспоминания, ностальгия, прошлое. Значение: Шестерка Кубков говорит о ностальгии и возвращении к воспоминаниям. Это карта прошлого и эмоциональной связи с ним.", "images/tarot/42.png"},
	{"Семерка Кубков", "Выбор, иллюзии, мечты. Значение: Семерка Кубков указывает на необходимость сделать выбор, предупреждая о возможных иллюзиях и фантазиях.", "images/tarot/43.png"},
	{"Восьмерка Кубков", "Отход, поиск, разочарование. Значение: Восьмерка Кубков говорит о поиске нового смысла и отдалении от старого. Это карта перемен и поиска себя.", "images/tarot/44.png"},
	{"Девятка Кубков", "Исполнение желаний, счастье, удовлетворение. Значение: Девятка Кубков — это карта исполнения желаний и чувства удовлетворения. Она указывает на завершение цикла и радость.", "images/tarot/45.png"},
	{"Десятка Кубков", "Счастье, гармония, семейное благополучие. Значение: Десятка Кубков символизирует семейное счастье и полную гармонию в отношениях.", "images/tarot/46.png"},
	{"Паж Кубков", "Романтика, воображение, творчество. Значение: Паж Кубков говорит о новых романтических чувствах и творческом подходе к жизни. Это карта эмоционального пробуждения.", "images/tarot/47.png"},
	{"Рыцарь Кубков", "Увлеченность, идеализм, поиски. Значение: Рыцарь Кубков символизирует стремление к идеалам и романтическим поискам. Это карта преданности своим чувствам.", "images/tarot/48.png"},
	{"Королева Кубков", "Забота, интуиция, сострадание. Значение: Королева Кубков — это карта заботы и сострадания. Она указывает на интуитивное понимание и мягкость в отношениях.", "images/tarot/49.png"},
	{"Король Кубков", "Эмоциональная зрелость, сочувствие, понимание. Значение: Король Кубков символизирует эмоциональную стабильность и зрелость. Это карта мудрости в чувствах и поддержки.", "images/tarot/50.png"},

	// Масть Мечей
	{"Туз Мечей", "Ясность, истина, идея. Значение: Туз Мечей указывает на новый интеллектуальный этап. Это карта истины и прояснения.", "images/tarot/51.png"},
	{"Двойка Мечей", "Выбор, сомнение, баланс. Значение: Двойка Мечей символизирует необходимость выбора, внутренний конфликт и поиск баланса между разумом и сердцем.", "images/tarot/52.png"},
	{"Тройка Мечей", "Боль, печаль, разбитое сердце. Значение: Тройка Мечей указывает на эмоциональную боль и разочарование. Это карта утраты и необходимости восстановления.", "images/tarot/53.png"},
	{"Четверка Мечей", "Отдых, восстановление, размышление. Значение: Четверка Мечей говорит о необходимости отдыха и размышлений. Это карта временного уединения и восстановления сил.", "images/tarot/54.png"},
	{"Пятерка Мечей", "Конфликт, поражение, жертва. Значение: Пятерка Мечей указывает на конфликт и борьбу, предупреждая о трудностях и возможных потерях.", "images/tarot/55.png"},
	{"Шестерка Мечей", "Перемены, переход, решение проблем. Значение: Шестерка Мечей символизирует переход на новый этап и необходимость оставить прошлое позади.", "images/tarot/56.png"},
	{"Семерка Мечей", "Обман, хитрость, тайные действия. Значение: Семерка Мечей указывает на необходимость быть осторожным с планами и избегать обмана.", "images/tarot/57.png"},
	{"Восьмерка Мечей", "Ограничения, страх, зависимость. Значение: Восьмерка Мечей символизирует чувство ограничения и страха, но напоминает о возможности освободиться от собственных страхов.", "images/tarot/58.png"},
	{"Девятка Мечей", "Тревога, страхи, бессонница. Значение: Девятка Мечей указывает на внутренние страхи и тревоги. Это карта борьбы с ментальными вызовами.", "images/tarot/59.png"},
	{"Десятка Мечей", "Завершение, поражение, боль. Значение: Десятка Мечей говорит о конце трудного периода и готовности к новому началу.", "images/tarot/60.png"},
	{"Паж Мечей", "Любознательность, анализ, наблюдательность. Значение: Паж Мечей символизирует желание узнавать и исследовать. Это карта открытости к новым знаниям.", "images/tarot/61.png"},
	{"Рыцарь Мечей", "Решительность, скорость, смелость. Значение: Рыцарь Мечей указывает на стремительное движение к цели и решительность в действиях.", "images/tarot/62.png"},
	{"Королева Мечей", "Интеллект, проницательность, честность. Значение: Королева Мечей символизирует рациональное мышление и способность к анализу.", "images/tarot/63.png"},
	{"Король Мечей", "Логика, лидерство, справедливость. Значение: Король Мечей — это карта логики, интеллекта и честности.", "images/tarot/64.png"},

	// Масть Пентаклей
	{"Туз Пентаклей", "Возможности, изобилие, процветание. Значение: Туз Пентаклей символизирует новые финансовые возможности и материальный успех. Это карта процветания и хороших начинаний.", "images/tarot/65.png"},
	{"Двойка Пентаклей", "Баланс, адаптация, многозадачность. Значение: Двойка Пентаклей говорит о необходимости находить баланс между различными аспектами жизни. Это карта гибкости и адаптации.", "images/tarot/66.png"},
	{"Тройка Пентаклей", "Сотрудничество, мастерство, работа в команде. Значение: Тройка Пентаклей символизирует важность командной работы и признания мастерства. Это карта совместных усилий для достижения целей.", "images/tarot/67.png"},
	{"Четверка Пентаклей", "Сдержанность, контроль, накопление. Значение: Четверка Пентаклей говорит о желании контролировать свои ресурсы и материальные блага. Это карта осторожности и сбережений.", "images/tarot/68.png"},
	{"Пятерка Пентаклей", "Потеря, трудности, нищета. Значение: Пятерка Пентаклей символизирует трудные времена и финансовые проблемы. Это карта необходимости поиска поддержки в трудные времена.", "images/tarot/69.png"},
	{"Шестерка Пентаклей", "Благотворительность, щедрость, поддержка. Значение: Шестерка Пентаклей говорит о важности делиться и поддерживать других. Это карта благотворительности и взаимопомощи.", "images/tarot/70.png"},
	{"Семерка Пентаклей", "Ожидание, усилия, оценка результатов. Значение: Семерка Пентаклей символизирует необходимость оценивать результаты своих усилий. Это карта терпения и анализа.", "images/tarot/71.png"},
	{"Восьмерка Пентаклей", "Труд, усердие, совершенствование. Значение: Восьмерка Пентаклей говорит о необходимости усердной работы и стремления к совершенству. Это карта роста и мастерства.", "images/tarot/72.png"},
	{"Девятка Пентаклей", "Самодостаточность, успех, материальное благополучие. Значение: Девятка Пентаклей символизирует независимость и достижения в материальном плане. Это карта комфорта и уверенности.", "images/tarot/73.png"},
	{"Десятка Пентаклей", "Богатство, наследие, семья. Значение: Десятка Пентаклей говорит о долгосрочных достижениях и наследии. Это карта семейных ценностей и материального процветания.", "images/tarot/74.png"},
	{"Паж Пентаклей", "Ученичество, желание учиться, практичность. Значение: Паж Пентаклей символизирует стремление к обучению и развитию. Это карта новых начинаний в учебе и карьере.", "images/tarot/75.png"},
	{"Рыцарь Пентаклей", "Стабильность, преданность, трудолюбие. Значение: Рыцарь Пентаклей говорит о надежности и упорном труде. Это карта преданности своим целям.", "images/tarot/76.png"},
	{"Королева Пентаклей", "Заботливость, практичность, изобилие. Значение: Королева Пентаклей символизирует заботливую и практичную природу. Это карта изобилия и мудрого управления ресурсами.", "images/tarot/77.png"},
	{"Король Пентаклей", "Материальное обеспечение, стабильность, лидерство. Значение: Король Пентаклей указывает на успешное управление и материальное благосостояние. Это карта уверенного лидерства в материальных вопросах.", "images/tarot/78.png"},
}

func GetRandomCard() TarotCard {
	return tarotCards[rand.Intn(len(tarotCards))]
}