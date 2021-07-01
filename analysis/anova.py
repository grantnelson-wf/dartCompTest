import pandas as pd
import statsmodels.api as sm
from statsmodels.formula.api import ols
import matplotlib.pyplot as plt
import seaborn as sns
import argparse;

# process arguments
parser = argparse.ArgumentParser(description='Perform an ANOVA on a space separarated input file.')
parser.add_argument('input', nargs='?', default='../results.txt', help='The input file to analyse (default is "results.txt").')
args = parser.parse_args()

# load data file
print('analysing: '+args.input)
df = pd.read_csv(args.input, sep=' ')

 # Perform 2-way ANOVA using order and index as variables and seconds as the dependent.
formula = 'seconds ~ order + index + order:index'
model = ols(formula, data=df).fit()
table = sm.stats.anova_lm(model, typ=2)
print(table)

# Box plot the treatments by name.
ax = sns.boxplot(x='name', y='seconds', data=df, color='#99c2a2')
ax = sns.swarmplot(x="name", y="seconds", data=df, color='#7d0013')
plt.show()

# Line plot order by seconds to confirm order's invariance.
sns.lineplot(data=df, x="order", y="seconds")
plt.show()
