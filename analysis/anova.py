import pandas as pd
import statsmodels.api as sm
from statsmodels.formula.api import ols
import matplotlib.pyplot as plt
import seaborn as sns

# load data file
df = pd.read_csv('../results.txt', sep=' ')

 # Perform 2-way ANOVA using order and index as variables and seconds as the dependent.
formula = 'seconds ~ order + index + order:index'
model = ols(formula, data=df).fit()
table = sm.stats.anova_lm(model, typ=2)
print(table)

# Box plot the treatments by name.
ax = sns.boxplot(x='name', y='seconds', data=df, color='#99c2a2')
ax = sns.swarmplot(x="name", y="seconds", data=df, color='#7d0013')
plt.show()
