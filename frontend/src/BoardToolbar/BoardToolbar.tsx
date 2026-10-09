import Icon from '../Icon'
import styles from './BoardToolbar.module.css'

type BoardToolbarProps = { cardCount: number; onAddColumn: () => void }

const BoardToolbar = ({ cardCount, onAddColumn }: BoardToolbarProps) => (
  <div className={styles.boardToolbar}>
    <div className={styles.viewTabs}>
      <button className={styles.viewTabActive} type="button"><Icon>▥</Icon> Board</button>
      <button className={styles.viewTab} type="button" title="List view coming soon"><Icon>☷</Icon> List</button>
    </div>
    <div className={styles.toolbarRight}>
      <span className={styles.taskCount}>{cardCount} {cardCount === 1 ? 'task' : 'tasks'}</span>
      <span className={styles.toolbarDivider} />
      <button className={styles.filterButton} type="button" title="Filters"><Icon>☷</Icon><span>Filter</span></button>
      <button className={styles.addColumnButton} type="button" onClick={onAddColumn}><span>＋</span> Add column</button>
    </div>
  </div>
)

export default BoardToolbar
