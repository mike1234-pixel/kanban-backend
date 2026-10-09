import type { Card, Column } from '../api'
import styles from './ColumnView.module.css'

type ColumnViewProps = {
  column: Column
  index: number
  draggingId: string
  onAdd: () => void
  onEdit: (card: Card) => void
  onDelete: (card: Card) => void
  onDragStart: (id: string) => void
  onDrop: (id: string) => void
}

const accents = [styles.accentLavender, styles.accentBlue, styles.accentAmber, styles.accentGreen]
const priorities = [styles.priorityCoral, styles.priorityGold, styles.priorityMint]
const priorityLabels = ['High priority', 'Medium priority', 'Low priority']
const assignees = ['Morgan T.', 'Jamie L.', 'Alex R.']

const ColumnView = ({ column, index, draggingId, onAdd, onEdit, onDelete, onDragStart, onDrop }: ColumnViewProps) => {
  const cards = [...(column.cards ?? [])].sort((a, b) => a.order - b.order)

  return (
    <section
      className={`${styles.column} ${draggingId ? styles.columnDropTarget : ''}`}
      onDragOver={(event) => event.preventDefault()}
      onDrop={(event) => {
        event.preventDefault()
        onDrop(event.dataTransfer.getData('text/plain'))
      }}
    >
      <div className={styles.columnHeading}>
        <div className={`${styles.columnAccent} ${accents[index % accents.length]}`} />
        <h2>{column.title}</h2>
        <span className={styles.columnCount}>{cards.length}</span>
        <button className={styles.columnMenu} aria-label={`Options for ${column.title}`} type="button">···</button>
      </div>

      <div className={styles.cardStack}>
        {cards.map((card, cardIndex) => {
          const metaIndex = (cardIndex + index) % 3
          return (
            <article
              key={card.id}
              className={`${styles.taskCard} ${draggingId === card.id ? styles.dragging : ''}`}
              draggable
              onDragStart={(event) => {
                event.dataTransfer.setData('text/plain', card.id)
                event.dataTransfer.effectAllowed = 'move'
                onDragStart(card.id)
              }}
              onDragEnd={() => onDragStart('')}
            >
              <div className={styles.cardMeta}>
                <span className={`${styles.priority} ${priorities[metaIndex]}`}><i />{priorityLabels[metaIndex]}</span>
                <button className={styles.cardMenu} type="button" aria-label={`Edit ${card.title}`} onClick={() => onEdit(card)}>···</button>
              </div>
              <button className={styles.cardTitle} type="button" onClick={() => onEdit(card)}>{card.title}</button>
              {card.description && <p className={styles.cardDescription}>{card.description}</p>}
              <div className={styles.cardFooter}>
                <div className={styles.cardAssignee}>
                  <span className={`${styles.assigneeAvatar} ${cardIndex % 2 ? styles.avatarRose : ''}`}>{['M', 'J', 'A'][cardIndex % 3]}</span>
                  <span>{assignees[cardIndex % 3]}</span>
                </div>
                <span className={styles.cardDate}>◷ &nbsp;{['Today', 'Tomorrow', 'Oct 24'][cardIndex % 3]}</span>
              </div>
              <button className={styles.deleteCard} type="button" onClick={() => onDelete(card)} aria-label={`Delete ${card.title}`}>×</button>
            </article>
          )
        })}
      </div>

      <button className={styles.addTask} type="button" onClick={onAdd}><span>＋</span> Add task</button>
    </section>
  )
}

export default ColumnView
